package crypto_test

import (
	"envkey/envkey-fetch/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPrivkey(t *testing.T) {
	_, err := crypto.ReadPrivkey(encryptedPrivkey, validPassphrase)
	assert.Nil(t, err, "Should not return an error.")
}

func TestEncrypt(t *testing.T) {
	decryptedPrivkey, _ := crypto.ReadPrivkey(rawEnvEncryptedPrivkey, rawEnvPassphrase)
	keys, _ := crypto.MakeKeyring(decryptedPrivkey, pubkeyArmored)
	_, err := crypto.Encrypt([]byte("test message"), keys)
	assert.Nil(t, err, "Should not return an error.")
}

func TestDecryptAndVerify(t *testing.T) {
	var err error
	decryptedPrivkey, _ := crypto.ReadPrivkey(rawEnvEncryptedPrivkey, rawEnvPassphrase)
	keys, _ := crypto.MakeKeyring(decryptedPrivkey, pubkeyArmored)
	_, err = crypto.DecryptAndVerify(signedEncryptedMessage, keys)
	assert.Nil(t, err, "Should not return an error.")

	keysWithInvalidPubkey, _ := crypto.MakeKeyring(decryptedPrivkey, invalidPubkeyArmored)
	_, err = crypto.DecryptAndVerify(signedEncryptedMessage, keysWithInvalidPubkey)
	assert.NotNil(t, err, "Should return an error.")
}

func TestVerifySignedCleartext(t *testing.T) {
	var err error

	pubkey, _ := crypto.ReadArmoredKey(pubkeyArmored)
	_, err = crypto.VerifySignedCleartext(signedMessage, pubkey)
	assert.Nil(t, err, "Should not return an error.")

	invalidPubkey, _ := crypto.ReadArmoredKey(invalidPubkeyArmored)
	_, err = crypto.VerifySignedCleartext(signedMessage, invalidPubkey)
	assert.NotNil(t, err, "Should return an error.")
}

func TestVerifyPubkeySignature(t *testing.T) {
	var err error

	pubkey, _ := crypto.ReadArmoredKey(pubkeyArmored)
	signedPubkey, _ := crypto.ReadArmoredKey(signedPubkeyArmored)
	invalidPubkey, _ := crypto.ReadArmoredKey(invalidPubkeyArmored)

	err = crypto.VerifyPubkeySignature(signedPubkey, pubkey)
	assert.Nil(t, err, "Should not return an error.")

	err = crypto.VerifyPubkeySignature(signedPubkey, invalidPubkey)
	assert.NotNil(t, err, "Should return an error.")
}

func TestVerifyPubkeyArmoredSignature(t *testing.T) {
	var err error
	err = crypto.VerifyPubkeyArmoredSignature(signedPubkeyArmored, pubkeyArmored)
	assert.Nil(t, err, "Should not return an error.")

	err = crypto.VerifyPubkeyArmoredSignature(signedPubkeyArmored, invalidPubkeyArmored)
	assert.NotNil(t, err, "Should return an error.")
}

func TestVerifyPubkeyWithPrivkey(t *testing.T) {
	var err error

	decryptedPrivkey, _ := crypto.ReadPrivkey(encryptedPrivkey, validPassphrase)
	pubkey, _ := crypto.ReadArmoredKey(pubkeyArmored)
	invalidPubkey, _ := crypto.ReadArmoredKey(invalidPubkeyArmored)

	err = crypto.VerifyPubkeyWithPrivkey(pubkey, decryptedPrivkey)
	assert.Nil(t, err, "Should not return an error.")

	err = crypto.VerifyPubkeyWithPrivkey(invalidPubkey, decryptedPrivkey)
	assert.NotNil(t, err, "Should return an error.")
}

var encryptedPrivkey = []byte("-----BEGIN PGP PRIVATE KEY BLOCK-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nxcMGBFmBC4IBCACl7nerYYcOByK8ytPwN1MUohF94c9Xkl9sF4upzKmioiA8\niNwYwcE7fzd3r5lsJg/Kfijf7kfa083okzHufHSwPWt6WSe7svPmVpq0g+qr\na8vHNFteDv9h1V1EAzCe9iB0BpsVwJ4eHv27cCzpkdu4G5jpoG+7LVS9DSe7\n332JwuzMGQBZ8yg7UxkhTHbIfVFBEh+Ae1OINiVaMGC52BBAS+1bIC0VRx+w\n9F+QX5mRaF4QVzffrGOlU6T04QmLCTTJvRA6kM1zxPNYNZB0XKRD/MiRg31A\nD5jqY3adLW6xPoe5Q6YSXpJfXTh8Nahsy3lzkc2gssUuGgOACr5HIupzABEB\nAAH+CQMIlxW2N1qoXnBgHSuCdTGZfVVzWlKRo+4yRNuTBRQY5PKMUTA5Gqjn\nc6rmxy8QUxOXqITxzozTyrL06i6yBjaEYSC+LQGOlB+8SjQxEo3PGBCkp2FJ\n75HXgUoSIEIZIhtn4qzgWUIJf9pzNMoGE2UggXgL7gQn8bE8qctPSOHWtP4c\nKugGJeToPxz/jMxXzGwpdkt69K1355Szl7z/cLXs1S686Xfzg5y5CF65zmoN\nk28ugTwFdEhaJJys3a2x9yz/ZnnR7bGFMXcotvJt+PdhQLy7GOhxtPLcBfE0\naqSSIHU+/lpL2+8Hs6uoGJmh0LaXwTOYGMsLDWsT9WjxqLkidbU4dKSYJvIJ\nrGyhlY5m9y0mGf7J9yEM+sYnH/ushB9lF0X+yGLXWoQiUvP+RxP1TRp9cGn7\nXdxXd28T0oWFX+WPGeWWK+avB2yHbgpmLK4bSSTIXtxpE9ahjOyW43ggBKk9\nOOPCXwMyMPZdxEVQeNRo5JdVxWhDyn9tFfbi03fkKCe2fh96RrbuyqPCd3cb\nMlpo4HwRSJrvWjtl3N+6W+gMMaFNzkEceRTt3g2lTx/aawzs2DukNYponERs\nT2wbGR2gVScbpRsQWHnYmT437KyuGIbr5oRLkIr/Rm0HJeR0W0d6swkjJNim\nVNuHvwh8h9Di7giTtV1yeJjDvsAMYy2+iWzgR4Sudh2IQfGDYAnt2Sh8eUJ0\nzJka0mu2RdiRFMKe4LfYimjpq7X2cw6DH2fj5Mv+tEEyMQ8k6RewmTfEVT9B\nurL/2ROXnnmgMyQTI7qhtGJx2eqTIGkW+4RV7/v3XB04WT/fpzmYMb8tR5mX\nEtQLjatwhl1Jdk0R6NdF7FPXm/O6emSA0MUpCnO6Nfo5WtvPwR9VTy8HW+2o\nEqNPb9+dSw8We/ucGaV9d0PfQJ5szUKEzY41YTViNGU3Y2JiYzMyY2Y3YmZj\nNDZkMmQxYmI2ZWNmZDkyNWNhOTVhNmQ3YjRkNzQ4NDAwZDFmYWRkNTNmZDEz\nIDw1YTViNGU3Y2JiYzMyY2Y3YmZjNDZkMmQxYmI2ZWNmZDkyNWNhOTVhNmQ3\nYjRkNzQ4NDAwZDFmYWRkNTNmZDEzQGVudmtleS5jb20+wsB1BBABCAApBQJZ\ngQuDBgsJBwgDAgkQTfvfGRGDNrQEFQgKAgMWAgECGQECGwMCHgEAAF1oCACH\n1O9UAa9WDwtxGECayulJauLZuRdpQs+O36XFlvYBYBz+WxHB2bah0QoYmG/8\n85wAI+GmE+rcf88RI2LJIKeu6EupQUUhIH81VNt9MtXxRkYxJIBD+GgvL7YA\nkGkddcUCeXBXrnEPQeIfxlhZ5eAN6X4tkp+Y1nbX480DkeRXWzgZZ4wOwCHo\ndTJ6+Rq5Y+G5VWbB5HUVhoK5kb5hWM+L2uQAsoNBdNgAPIpU60jqIgP0O2wZ\nnLjaBaaE1iJI1aIdLvZkHXco7KmXM+U8lB33+YsodIwPnHUDWWW6lMuhmO+r\nWNH7QaoHTJ2IZeatQ9cix3jyJGUHE3AlXPa1VhR/x8MGBFmBC4IBCADA/wLD\n7UvzZDx/IOiKtmGRpXl5Ajw85u9ju3yuNb96lsR1hGiR6j1FklbrEa9eZAXn\n+Qh8keb20Pw5a7qwtl5A8WV0ycAsHQCUuKZayoKksbwINbgRQPFieqJeuAPH\nurX4nSdwSlyabkxRLyXzIGSYRlwKX6b8I69eKtSYLLIlWQdOTe+hOREOgNKm\nHqrmCzhJb9fjqixe9vuSbb35sThOwxZUslxbYWtC/5h8fmBfqY3SsQDEhglP\nDILXcvFYYaxPI6XQDMFb6xGjPf7A6Z0/Azau5OzMiwMfaOUnLnep54/wKXxG\nUPqqZVA6QnPqv4xVBa//THMKcN49Rnk2EpfhABEBAAH+CQMI6W6PcCtS7aZg\nBI/ILU27ZRtmtj+fBEC3qBcnqrOsQ56GLffWe+Jy6kANpt1tQmOiO6CBBc3i\nSVj5ASdw1HuhsAWyeYnOAlPfy60qVj+eWJmvWRA/s8Rzax31Wjhu5z+Zg6Bu\nSAVoUIb3XoxuBtadgKJ7CV1XG+zVxGpIf5lJviQ2IrANIpouoVbAmzJJ14hQ\nPQUT39YL01L+OyT8HIXQh55pFI4nQBTNsA1dhkXbOVqMbGtzPfA9pdiJxbK9\nfVJLMW7UlrynkTZMdgdpJNS+cuf4EhLQtQCQrQ3rfHSPTRvUMmtMY9rj7MCE\n+XKlpxo2940R7zuZLqj58M4aVuLBFhoRDswbI/mwaDwdpGP/BUYagYmu8hYe\nWpEU+/djdxoWWrEzzI1gTF7OcNP2/QmsnKAxjsfISgR/tJQJTB8ig85Hrtpw\nORyN1bp+S4wo572Bd2FHP26K2bdsV0VKbCL+4oihhiauYfVv2xM3+nh39VHr\nSPix4aR9rPGiSq9gn7/zkcN67Fk2qHzlo8jb6BGWpBTOjEow8AzCEMHBliZR\nA/EeUwzS+9AiVtADoGPM4pmKJ0uhdF8lWoRGgEUQyytHAWgd5CUf7+Ih8AsX\npETaPbbIQ7KoxevAkRNLKlbcqY4OPeoLRjhnsNVRCD6BAYKZDFN4E0aAG1jQ\nl+vjfWKD7UKT0uqvy1Fy2T6vQ4dvxkezNPIGX3s2FTN3t3WTqv+Zm6VNNGRC\np9nUkN+5Gk+wvMy9wJ5vroKTBpdOHnMmfqCSqU5OsjV1a39vH841P31valv5\npR6PQz+0iTWL6BFPh2U5GrSUFd/pU7f7HDMlxX+RrmJD2iYPjrhdBvAA/dZF\nHd9ptagYYeHG/Ue7TnwAEon6GbV7LMd7Fub2a0IZgCIWu+tDPqG/zC6KnYDc\nZsOPLdxYmUKawsBfBBgBCAATBQJZgQuDCRBN+98ZEYM2tAIbDAAAiPgH/iIv\nMBUMOkSDMyMlY60m3nppVwJ12us0GmiReQ23oKyLTExJQsMUn4bC5Z7gULrP\nY6p90JlIpIrS6ufYKNW2OYow62LqfbMezIiM9rp8A2I3kBPaM3bj+0UFgarR\nN6k9qr1Oonino+Xhf5BnHRdTh1O4LONWY8GC3kISbpd79kvw5QI6bstc3lVX\npm8YDhzOpi/zwj5r93trSRLDEB9cUd0PLBKhSJN2ziC420uhi9fx7LwNG1kw\nLxa2qV9NYUv5Q0p6VOzb8C43YR72hs3KUXjqBe7ehYXN8RjJ4az0jSLpkRru\nU0kCHQNhECPOmEiaENQVehzH93e96Vn+gK8YCEw=\r\n=EDGr\r\n-----END PGP PRIVATE KEY BLOCK-----\r\n")

var validPassphrase = []byte("passworded")

var pubkeyArmored = []byte("-----BEGIN PGP PUBLIC KEY BLOCK-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nxsBNBFmBC4IBCACl7nerYYcOByK8ytPwN1MUohF94c9Xkl9sF4upzKmioiA8\niNwYwcE7fzd3r5lsJg/Kfijf7kfa083okzHufHSwPWt6WSe7svPmVpq0g+qr\na8vHNFteDv9h1V1EAzCe9iB0BpsVwJ4eHv27cCzpkdu4G5jpoG+7LVS9DSe7\n332JwuzMGQBZ8yg7UxkhTHbIfVFBEh+Ae1OINiVaMGC52BBAS+1bIC0VRx+w\n9F+QX5mRaF4QVzffrGOlU6T04QmLCTTJvRA6kM1zxPNYNZB0XKRD/MiRg31A\nD5jqY3adLW6xPoe5Q6YSXpJfXTh8Nahsy3lzkc2gssUuGgOACr5HIupzABEB\nAAHNjjVhNWI0ZTdjYmJjMzJjZjdiZmM0NmQyZDFiYjZlY2ZkOTI1Y2E5NWE2\nZDdiNGQ3NDg0MDBkMWZhZGQ1M2ZkMTMgPDVhNWI0ZTdjYmJjMzJjZjdiZmM0\nNmQyZDFiYjZlY2ZkOTI1Y2E5NWE2ZDdiNGQ3NDg0MDBkMWZhZGQ1M2ZkMTNA\nZW52a2V5LmNvbT7CwHUEEAEIACkFAlmBC4MGCwkHCAMCCRBN+98ZEYM2tAQV\nCAoCAxYCAQIZAQIbAwIeAQAAXWgIAIfU71QBr1YPC3EYQJrK6Ulq4tm5F2lC\nz47fpcWW9gFgHP5bEcHZtqHRChiYb/zznAAj4aYT6tx/zxEjYskgp67oS6lB\nRSEgfzVU230y1fFGRjEkgEP4aC8vtgCQaR11xQJ5cFeucQ9B4h/GWFnl4A3p\nfi2Sn5jWdtfjzQOR5FdbOBlnjA7AIeh1Mnr5Grlj4blVZsHkdRWGgrmRvmFY\nz4va5ACyg0F02AA8ilTrSOoiA/Q7bBmcuNoFpoTWIkjVoh0u9mQddyjsqZcz\n5TyUHff5iyh0jA+cdQNZZbqUy6GY76tY0ftBqgdMnYhl5q1D1yLHePIkZQcT\ncCVc9rVWFH/OwE0EWYELggEIAMD/AsPtS/NkPH8g6Iq2YZGleXkCPDzm72O7\nfK41v3qWxHWEaJHqPUWSVusRr15kBef5CHyR5vbQ/DlrurC2XkDxZXTJwCwd\nAJS4plrKgqSxvAg1uBFA8WJ6ol64A8e6tfidJ3BKXJpuTFEvJfMgZJhGXApf\npvwjr14q1JgssiVZB05N76E5EQ6A0qYequYLOElv1+OqLF72+5JtvfmxOE7D\nFlSyXFtha0L/mHx+YF+pjdKxAMSGCU8Mgtdy8VhhrE8jpdAMwVvrEaM9/sDp\nnT8DNq7k7MyLAx9o5Scud6nnj/ApfEZQ+qplUDpCc+q/jFUFr/9Mcwpw3j1G\neTYSl+EAEQEAAcLAXwQYAQgAEwUCWYELgwkQTfvfGRGDNrQCGwwAAIj4B/4i\nLzAVDDpEgzMjJWOtJt56aVcCddrrNBpokXkNt6Csi0xMSULDFJ+GwuWe4FC6\nz2OqfdCZSKSK0urn2CjVtjmKMOti6n2zHsyIjPa6fANiN5AT2jN24/tFBYGq\n0TepPaq9TqJ4p6Pl4X+QZx0XU4dTuCzjVmPBgt5CEm6Xe/ZL8OUCOm7LXN5V\nV6ZvGA4czqYv88I+a/d7a0kSwxAfXFHdDywSoUiTds4guNtLoYvX8ey8DRtZ\nMC8WtqlfTWFL+UNKelTs2/AuN2Ee9obNylF46gXu3oWFzfEYyeGs9I0i6ZEa\n7lNJAh0DYRAjzphImhDUFXocx/d3velZ/oCvGAhM\r\n=w9ow\r\n-----END PGP PUBLIC KEY BLOCK-----\r\n\r\n")

var invalidPubkeyArmored = []byte("-----BEGIN PGP PUBLIC KEY BLOCK-----\r\nVersion: OpenPGP.js v2.3.6\r\nComment: http://openpgpjs.org\r\n\r\nxsBNBFi54G8BCACjHLBx/K+JSrjYBWO8o2yPi3htkDPBGe2B/wZDyuHlVKCF\nViYdIfp+2phN6dyqd/Ts7FheKAzBvBVtJq2UCMrFhuP8Ni1Obx9YEtiN3lwh\nCCG9UmBi8Up19Siz80fqdAJyMFKbi46P2RuZwFpViaUDszwumLI27qh+4yKR\nq+KDweH4BQcS9ier6Ixy4ETGXXxjE2wzFrbX0zWJyoXtX+ksTqATC2y6AcO3\nExaso/IzgbbLuHuk5TkIg/8cQeIgbnT/eNrnGHefqZ/j/TsRz5Vx79sQeyW3\nxg7n7v78JkNyQGQvUEMCg2sUKcKRgKvMU+KxSScVGZby/Vpht4n3CIfZABEB\nAAHNjmEzZDI2NDBjYTYzNWE2Yzk3YjBlYWIwYmYxZDJkNmM3YTAyOWI0NTY3\nOGM3ODRhYTZiZGEwMjJkYmZiZjRhYmYgPGEzZDI2NDBjYTYzNWE2Yzk3YjBl\nYWIwYmYxZDJkNmM3YTAyOWI0NTY3OGM3ODRhYTZiZGEwMjJkYmZiZjRhYmZA\nZW52a2V5LmNvbT7CwHUEEAEIACkFAli54G8GCwkHCAMCCRBGutAa65RGWQQV\nCAIKAxYCAQIZAQIbAwIeAQAA5mIH/iQr5Q1o9SKYs2FpivZlSAyMd75MGros\nZV0xPrDcQr08QeVV2HYIPDC8EjAFdZ1MK5Kr5XJpVQL94k8b0ELfEO7mloRk\nNl6xbSsGbJafR/K33+H/Hra6t7anDx/0gAC2xrGlLpICtibU77avqnIv8/Rs\ngOrwXHTCWLHtLBiil4IgHsMoJTAMMIafRPwL0r6P/pjKWyoHVLfYXeb5wr1I\noSoDuReMH2uklDhoOvzzaOEZnLCpZM3R5Iv0614ZD5F8HlC0F9fZ+kzcroNn\n1UVUG7UQpvA+JwYN+uOL7h/oYWm2BptGQNhtHc32uE6V2wCzvD3zLRQF47dv\nkp54d1DN7S3OwE0EWLngbwEIAM3StNaTB5FKc1wp8Paz5IyRaMtvy4pQeMJr\nyL3jLcG8+GQlTqvIPtrxzpLb6DtGiQeBpsCgVQO/SInrSgBV44Ojvztp8bLA\nLMLFtE7gpI9G7Ezr4n4he1lIZpzB5qk1XXtiKaIZSnRFuYUElFWsfaY2J9kL\n+SdapcXV7teQByeiWF1sMCW947lUoytxuQRu+x8zrhqsgLg3w0YyO+7YfC1Z\nu1AMbngInd7fccTJsD98m6kLn4K0DWsnBN42K1+t2l2qOWtnQ+a8GPzN6aml\nq2TcjQUr/JgQMIxfhMhTDECePI1MGCHhw6EmKyyCHmkzffzJyRNc32mBmF7V\nHsvwwBkAEQEAAcLAXwQYAQgAEwUCWLngcAkQRrrQGuuURlkCGwwAAPYXB/98\nKnfwgQvy0pBr+BQiH/k8iSmmZlql90GtY0KUvCVOwvfAEGpMyVLq+RkthFZT\nrliiYmfItzAUtLaD0xEJZqFYCnIhfDOlITrRKvr+yOODd5WQey3IiBjuN3RQ\nfNDGdLK6zbTc8lpQK7m/DTZ3xX+PGxXQGKtsVXLLWIZ6vZwK2Xb56XgwzRNU\nqovXWWyPIJccg7yLYe4X66uDA79mWlhdijmdlRxqJHkvz3Q7tZEkO5bt9fjc\nFciIAUoM1+klPATaInPeXqIrY5zCe5HDQFxvYURFC19tvJZRIjynyKk/Jco5\nS3G887Vd7NCDKpYHMbV7XGva7zBBTexzAPUt10l2\r\n=+YDH\r\n-----END PGP PUBLIC KEY BLOCK-----\r\n\r\n")

var signedPubkeyArmored = []byte("-----BEGIN PGP PUBLIC KEY BLOCK-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nxsBNBFmBEtUBCACgY9ybvYbW6fNhGCmUmoWaDZTn7CFYC8YRuKBfvmPTms1/\nP80VtY+hl9DYDMJu7V7iP8AwDWVNioqS3fW34RfFkF5/bVZfUEYUOoor0JHc\nHkIJgkayFNSpgiDucWIZ8TLcI/smUjpcT3epvR+tmBil7E0bY80EDEsYz3cd\nIdZGVV2yihwTSRqhPFVYD0iCWwqOAwuecFnpi6roUDXPHUUAuSbazwFsLTyg\nppunqya7le6DaeHX1bG8YPWNk21uK9oehZQbKAE1I7oBOG20vyFCZvStZAIe\nasGGnhE58u94fIv9aopiAz0ngCOyZ41YTp9KY+ZDOdvGkXPIvybIrqCXABEB\nAAHNjjAwNDljMDM3MzViNTEyMzRkNzg5YTA5ZDI1MGFiZWJjODVmMTk5YTNi\nOTBiNDc1NjZlZjY4YjI4ZDk0NDVhYzEgPDAwNDljMDM3MzViNTEyMzRkNzg5\nYTA5ZDI1MGFiZWJjODVmMTk5YTNiOTBiNDc1NjZlZjY4YjI4ZDk0NDVhYzFA\nZW52a2V5LmNvbT7CwHUEEAEIACkFAlmBEtYGCwkHCAMCCRAK77Qge/sL8gQV\nCAoCAxYCAQIZAQIbAwIeAQAA+UUIAIiPMR6ePRiytkJZ1arnmFHuunBgVr+I\n+g02OS6NK8VH2p3ekzud3rrTuMDqvwr9UjD7bb2tnnr9QKoATdsCcQNag3hE\ndCyreYH4r0xYSq5GorhgnGUosaStscHxBOZD76v9Ah5WHpqNh28vdvVUEAdD\nTYHqyALPYYgDJuvv5qrgreEnJavHZohwGOX9PQwdRR5bLbxLciXyf/F4O6Bb\nCEJaWqd4+WrLMKry8C4JafdLYUFlwkfP15SCemuJYCFORbb9qoXNBxWpH50S\nWaB0P4XDdHk36JHmeuOt/FSont7yL5GvXZjJBl5HNWh1TPV2mn5AlKJvh4B1\nTixr7OdDlbHCwF8EEAEIABMFAlmBEtcJEE373xkRgza0AhsDAAAWmggAlDN9\nOXvTz5bcAYoblPLJ97vP9Y4/ieGXH1AYfWQZO9GDIit/n1IgCJEsYxswnO5a\nb2MWEP2pzlQyNHlShN0x/JT7kND4ftLGImxfegmYgRBs50g2DKAu6/vs3Uzv\nqLQ2YoiXylI79pf20Dm/gemoPWTwDPlEgnBzniLfJXat6dM/Hp7VsnsZnWc5\nw3VEmvDPqiwFDDQyCqKYukPGA1lItRaWnwigLHdf6P3IwXKdPYGG1spvG2im\n6BueKkPuys73ItJ2eaJlyTw2YFK1CE6QnJnDv5ZFjAF/PSMFg1XyvAz1tHFz\nvlBdc51H4CUtcXucZvjR1QgbXvP6vT9Xy4iDas7ATQRZgRLVAQgAx5/blkUU\n5bNlFWJ/VtfZioAqpQv+U9RJI0gcChlom5KNVC/XNOApfJilR9PUxs2NZHXm\naQ8QnP2hYYAmAkC8Vq+FLZuPEzYnQPTDIRRoiyIHwRFlIAef8D+ynwyg1CGD\numG3g1WismK5cMoHdbFkvZxhTRYUrJ6s1lXe7Ycb9emD7l++HS5CaIzqnp2H\nfsdlcygJc0lFhJrr1fogpAxaUMFM3uuB5Pg03/az3BqEh7Nhsd9/BwINHHYJ\nnKTr1COw+XzhwzoRzaWt+MvAqoZ+90S41GmqGfV6s3g6f+Ou9PChjgC1ipGt\npirGNxjPnWlmoytQik6BJbWiHCBYRYWtFQARAQABwsBfBBgBCAATBQJZgRLW\nCRAK77Qge/sL8gIbDAAAfBwH/17bcGnFDcGb409wekzja51CRW/E3s+80l/4\nBLrC+qXp/eO/cdN1ymgzqsswBIE2Ror51rhmeTGw67lG6KJYGuoY524Iycwx\nB0G80phGqgrTkvhIyH5QH9wmm8I53slu1HTZRd5AdF6/WUxbafHY5NJlpYEb\nv9AhnGZ2GSBJQ9Q9hQb0EHzG/52pMO/1rpIraG07xlOXwM64u5nwhxLEmQj9\nz9E9xmfl9wZxqzEExtWcinblj7ZCtcdgOfLqHMH0ZGSBGg0bMr5tOjJ8rT3Z\nXtsv+1BXdAHMylzomWlU6fcog2t4uRKTTq5CTBzt5spsgncTR6awXuQ1Fro1\nmzY7//s=\r\n=mmUg\r\n-----END PGP PUBLIC KEY BLOCK-----\r\n\r\n")

var signedMessage = []byte("\r\n-----BEGIN PGP SIGNED MESSAGE-----\r\nHash: SHA256\r\n\r\n{\"7c3f7524-91d6-4d0a-b2db-9cf5ac4b9b89\":{\"type\":\"user\",\"pubkey\":\"-----BEGIN PGP PUBLIC KEY BLOCK-----\\r\\nVersion: OpenPGP.js v2.5.4\\r\\nComment: http://openpgpjs.org\\r\\n\\r\\nxsBNBFmBC4IBCACl7nerYYcOByK8ytPwN1MUohF94c9Xkl9sF4upzKmioiA8\\niNwYwcE7fzd3r5lsJg/Kfijf7kfa083okzHufHSwPWt6WSe7svPmVpq0g+qr\\na8vHNFteDv9h1V1EAzCe9iB0BpsVwJ4eHv27cCzpkdu4G5jpoG+7LVS9DSe7\\n332JwuzMGQBZ8yg7UxkhTHbIfVFBEh+Ae1OINiVaMGC52BBAS+1bIC0VRx+w\\n9F+QX5mRaF4QVzffrGOlU6T04QmLCTTJvRA6kM1zxPNYNZB0XKRD/MiRg31A\\nD5jqY3adLW6xPoe5Q6YSXpJfXTh8Nahsy3lzkc2gssUuGgOACr5HIupzABEB\\nAAHNjjVhNWI0ZTdjYmJjMzJjZjdiZmM0NmQyZDFiYjZlY2ZkOTI1Y2E5NWE2\\nZDdiNGQ3NDg0MDBkMWZhZGQ1M2ZkMTMgPDVhNWI0ZTdjYmJjMzJjZjdiZmM0\\nNmQyZDFiYjZlY2ZkOTI1Y2E5NWE2ZDdiNGQ3NDg0MDBkMWZhZGQ1M2ZkMTNA\\nZW52a2V5LmNvbT7CwHUEEAEIACkFAlmBC4MGCwkHCAMCCRBN+98ZEYM2tAQV\\nCAoCAxYCAQIZAQIbAwIeAQAAXWgIAIfU71QBr1YPC3EYQJrK6Ulq4tm5F2lC\\nz47fpcWW9gFgHP5bEcHZtqHRChiYb/zznAAj4aYT6tx/zxEjYskgp67oS6lB\\nRSEgfzVU230y1fFGRjEkgEP4aC8vtgCQaR11xQJ5cFeucQ9B4h/GWFnl4A3p\\nfi2Sn5jWdtfjzQOR5FdbOBlnjA7AIeh1Mnr5Grlj4blVZsHkdRWGgrmRvmFY\\nz4va5ACyg0F02AA8ilTrSOoiA/Q7bBmcuNoFpoTWIkjVoh0u9mQddyjsqZcz\\n5TyUHff5iyh0jA+cdQNZZbqUy6GY76tY0ftBqgdMnYhl5q1D1yLHePIkZQcT\\ncCVc9rVWFH/OwE0EWYELggEIAMD/AsPtS/NkPH8g6Iq2YZGleXkCPDzm72O7\\nfK41v3qWxHWEaJHqPUWSVusRr15kBef5CHyR5vbQ/DlrurC2XkDxZXTJwCwd\\nAJS4plrKgqSxvAg1uBFA8WJ6ol64A8e6tfidJ3BKXJpuTFEvJfMgZJhGXApf\\npvwjr14q1JgssiVZB05N76E5EQ6A0qYequYLOElv1+OqLF72+5JtvfmxOE7D\\nFlSyXFtha0L/mHx+YF+pjdKxAMSGCU8Mgtdy8VhhrE8jpdAMwVvrEaM9/sDp\\nnT8DNq7k7MyLAx9o5Scud6nnj/ApfEZQ+qplUDpCc+q/jFUFr/9Mcwpw3j1G\\neTYSl+EAEQEAAcLAXwQYAQgAEwUCWYELgwkQTfvfGRGDNrQCGwwAAIj4B/4i\\nLzAVDDpEgzMjJWOtJt56aVcCddrrNBpokXkNt6Csi0xMSULDFJ+GwuWe4FC6\\nz2OqfdCZSKSK0urn2CjVtjmKMOti6n2zHsyIjPa6fANiN5AT2jN24/tFBYGq\\n0TepPaq9TqJ4p6Pl4X+QZx0XU4dTuCzjVmPBgt5CEm6Xe/ZL8OUCOm7LXN5V\\nV6ZvGA4czqYv88I+a/d7a0kSwxAfXFHdDywSoUiTds4guNtLoYvX8ey8DRtZ\\nMC8WtqlfTWFL+UNKelTs2/AuN2Ee9obNylF46gXu3oWFzfEYyeGs9I0i6ZEa\\n7lNJAh0DYRAjzphImhDUFXocx/d3velZ/oCvGAhM\\r\\n=w9ow\\r\\n-----END PGP PUBLIC KEY BLOCK-----\\r\\n\\r\\n\",\"pubkeyFingerprint\":\"ad80ac3bcec7047db976a15a4dfbdf19118336b4\",\"invitedById\":null,\"invitePubkey\":null,\"invitePubkeyFingerprint\":null,\"email\":\"o@v50.com\",\"role\":\"org_owner\"},\"37645d91-566e-4a2b-b0d2-ef96a95c9f44\":{\"type\":\"user\",\"pubkey\":null,\"pubkeyFingerprint\":null,\"invitedById\":\"7c3f7524-91d6-4d0a-b2db-9cf5ac4b9b89\",\"invitePubkey\":\"-----BEGIN PGP PUBLIC KEY BLOCK-----\\r\\nVersion: OpenPGP.js v2.5.4\\r\\nComment: http://openpgpjs.org\\r\\n\\r\\nxsBNBFmBEtUBCACgY9ybvYbW6fNhGCmUmoWaDZTn7CFYC8YRuKBfvmPTms1/\\nP80VtY+hl9DYDMJu7V7iP8AwDWVNioqS3fW34RfFkF5/bVZfUEYUOoor0JHc\\nHkIJgkayFNSpgiDucWIZ8TLcI/smUjpcT3epvR+tmBil7E0bY80EDEsYz3cd\\nIdZGVV2yihwTSRqhPFVYD0iCWwqOAwuecFnpi6roUDXPHUUAuSbazwFsLTyg\\nppunqya7le6DaeHX1bG8YPWNk21uK9oehZQbKAE1I7oBOG20vyFCZvStZAIe\\nasGGnhE58u94fIv9aopiAz0ngCOyZ41YTp9KY+ZDOdvGkXPIvybIrqCXABEB\\nAAHNjjAwNDljMDM3MzViNTEyMzRkNzg5YTA5ZDI1MGFiZWJjODVmMTk5YTNi\\nOTBiNDc1NjZlZjY4YjI4ZDk0NDVhYzEgPDAwNDljMDM3MzViNTEyMzRkNzg5\\nYTA5ZDI1MGFiZWJjODVmMTk5YTNiOTBiNDc1NjZlZjY4YjI4ZDk0NDVhYzFA\\nZW52a2V5LmNvbT7CwHUEEAEIACkFAlmBEtYGCwkHCAMCCRAK77Qge/sL8gQV\\nCAoCAxYCAQIZAQIbAwIeAQAA+UUIAIiPMR6ePRiytkJZ1arnmFHuunBgVr+I\\n+g02OS6NK8VH2p3ekzud3rrTuMDqvwr9UjD7bb2tnnr9QKoATdsCcQNag3hE\\ndCyreYH4r0xYSq5GorhgnGUosaStscHxBOZD76v9Ah5WHpqNh28vdvVUEAdD\\nTYHqyALPYYgDJuvv5qrgreEnJavHZohwGOX9PQwdRR5bLbxLciXyf/F4O6Bb\\nCEJaWqd4+WrLMKry8C4JafdLYUFlwkfP15SCemuJYCFORbb9qoXNBxWpH50S\\nWaB0P4XDdHk36JHmeuOt/FSont7yL5GvXZjJBl5HNWh1TPV2mn5AlKJvh4B1\\nTixr7OdDlbHCwF8EEAEIABMFAlmBEtcJEE373xkRgza0AhsDAAAWmggAlDN9\\nOXvTz5bcAYoblPLJ97vP9Y4/ieGXH1AYfWQZO9GDIit/n1IgCJEsYxswnO5a\\nb2MWEP2pzlQyNHlShN0x/JT7kND4ftLGImxfegmYgRBs50g2DKAu6/vs3Uzv\\nqLQ2YoiXylI79pf20Dm/gemoPWTwDPlEgnBzniLfJXat6dM/Hp7VsnsZnWc5\\nw3VEmvDPqiwFDDQyCqKYukPGA1lItRaWnwigLHdf6P3IwXKdPYGG1spvG2im\\n6BueKkPuys73ItJ2eaJlyTw2YFK1CE6QnJnDv5ZFjAF/PSMFg1XyvAz1tHFz\\nvlBdc51H4CUtcXucZvjR1QgbXvP6vT9Xy4iDas7ATQRZgRLVAQgAx5/blkUU\\n5bNlFWJ/VtfZioAqpQv+U9RJI0gcChlom5KNVC/XNOApfJilR9PUxs2NZHXm\\naQ8QnP2hYYAmAkC8Vq+FLZuPEzYnQPTDIRRoiyIHwRFlIAef8D+ynwyg1CGD\\numG3g1WismK5cMoHdbFkvZxhTRYUrJ6s1lXe7Ycb9emD7l++HS5CaIzqnp2H\\nfsdlcygJc0lFhJrr1fogpAxaUMFM3uuB5Pg03/az3BqEh7Nhsd9/BwINHHYJ\\nnKTr1COw+XzhwzoRzaWt+MvAqoZ+90S41GmqGfV6s3g6f+Ou9PChjgC1ipGt\\npirGNxjPnWlmoytQik6BJbWiHCBYRYWtFQARAQABwsBfBBgBCAATBQJZgRLW\\nCRAK77Qge/sL8gIbDAAAfBwH/17bcGnFDcGb409wekzja51CRW/E3s+80l/4\\nBLrC+qXp/eO/cdN1ymgzqsswBIE2Ror51rhmeTGw67lG6KJYGuoY524Iycwx\\nB0G80phGqgrTkvhIyH5QH9wmm8I53slu1HTZRd5AdF6/WUxbafHY5NJlpYEb\\nv9AhnGZ2GSBJQ9Q9hQb0EHzG/52pMO/1rpIraG07xlOXwM64u5nwhxLEmQj9\\nz9E9xmfl9wZxqzEExtWcinblj7ZCtcdgOfLqHMH0ZGSBGg0bMr5tOjJ8rT3Z\\nXtsv+1BXdAHMylzomWlU6fcog2t4uRKTTq5CTBzt5spsgncTR6awXuQ1Fro1\\nmzY7//s=\\r\\n=mmUg\\r\\n-----END PGP PUBLIC KEY BLOCK-----\\r\\n\\r\\n\",\"invitePubkeyFingerprint\":\"ba420ae16a58122180d07a4a0aefb4207bfb0bf2\",\"email\":\"a1@v50.com\",\"role\":\"basic\"}}\r\n-----BEGIN PGP SIGNATURE-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nwsBcBAEBCAAQBQJZgRLaCRBN+98ZEYM2tAAA7a0H/A72rWfLpgDumnF7dpDL\nlO0ODeKeDt6q8OLG6teQEFf5+gJMGLkBUwgNwHKmKI907EgB/HrjS2BtkIbC\nOgDlbirP5WXXkhqv6cbnV1Y1zI7yOZJYlBehhO5vGElffR5ujaNykqtdn23F\nItaYgQGvE8c40rWiXw2F4zWJG8Ui/T5kd/YWQmJrjI5UyS7nxPg7y/HfokOs\nI174tAqwesEQWykRNEGmQE/FGea/m5UHTz+hDx1zvC0cL1aCt1gU2qQl0gMh\nn/rSAqN4xzg3HyJ1FKBHfbQzYWmhYA4WmXxwLwMPukiDBwkHgpQ2NoI9QUpa\nfz2vZq2ATVNC3rFAAOBhUjU=\r\n=gBxS\r\n-----END PGP SIGNATURE-----\r\n")

var signedEncryptedMessage = []byte("-----BEGIN PGP MESSAGE-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nwcBMA12U1SdLsUmiAQgAheOouojhwHHibVWrsZW/s4N3+2VK7BG4wfrqQBd2\nSsEP67V1YVo3o8soGXy/bwR7HeEDOsl3+oeAYhnQkPa9BiBFSDvjFIQkKXYr\nS/4A4bJiO5dt3Fzb6iBsIIwFlvyIi5RLASQyNBST2CiM8zVd1Kfh5Ty3nAAv\nHGUxPq2Rh1PzsOZPvJ5DvyErBkc5crexRsDtSGT6dxMF2l6iZAZQzYPttYWl\nwKoCtaCNj09oLNcAeiUWe2WqpawyT/ARPNAOdgYOlAZ8r9/g3+DoqpAe7gRu\ntuVS6XGkMdeZCaqaKWW0mHWkC8vuCCauOT89lXuQorFXZDA4lViRvEsx8F6T\nS9LAywEaq0UrM59ctEGwZuw5V4b6+Gg4TW+Lrenra2rknpqqhWHoJ9ot2d1a\nMpZBPaZkVlCcCuGbMuvkFogf3PD0NY/S3yEn85UW/P5IaTnrI08PiPTDAj2i\nQaFvctO6eo0vOELZX61rjzaclRNOFD53jdOjprC+ZaUDsJzsFPUrtKCUCfrJ\nd/GAzm9GKBkEprXd7xbtyft59uRMfP2SZxocQx61PayaQIUHTs4pDK4BRSOd\nbfJLh2g9OYQNXMIJvsTDAM7MKiTccsYZi1e/W2UUg8TblGxS6+P1gzBrxXcU\nEIuWGRy+I5RNhigFsS8xvwcylryZOwvJy+S4sOu50I0lNFlbUMqGhg32znRv\njp1bGzGGMsg23G3EiYDeJjNsI+fJ9zinWqoYbXdDMEYVzpARF2DxlSrDmi6g\nzGSDdAXX0vk3aSdzCLCxX9ZH8npOoKN3DTmAFnznZ3Sg8TLjuerHcqYldKPP\nQXYzNn0bMC5qAzFF7OdoBrek18Eb+jbtKuTksAbe2lSy1kij30Ft\r\n=efYw\r\n-----END PGP MESSAGE-----\r\n")

var rawEnvEncryptedPrivkey = []byte("-----BEGIN PGP PRIVATE KEY BLOCK-----\r\nVersion: OpenPGP.js v2.5.4\r\nComment: http://openpgpjs.org\r\n\r\nxcMGBFmBFecBCADlKkhP4WfTNrd5sIzk+zJtND4R+W7cDZUkCOwmxTEdnuXR\nGUDOH3c9iNVBOVYtwQ77wDRaQCucwPrKoVUfJcOASg/McyGhcswD0N4oB5pN\nWTK0FMCfUsz/njznmw1sbPlNNmKNOWhC8HGdL/cwf5gjp1JvFDZh6mPfxYx7\nZxcYZonaTcd4FykgW7+zl2M9EaRJlSs/snE+8AvoGK5Rm5Dk5+LSRd4eBwm2\nlBV54NnHh9qhoI4HBDT2a2DRlIwHD0e8MJ4G/vfgM6ThJd/Zjyy+csdiqlfr\n5awi6K/ojP0oTCbEVzYuq1GVl+cz/CV9AdQADD/EThk2/uHfa4xlcdvFABEB\nAAH+CQMIoW77jxv08CtgmMpw5hnDQzrZ2/pNqWYcgKwWwhnmaIS1HMoyetoa\nO+6MgHi69eW+TAZZivBrMYqmBLXCJBdxbGX6rSUs94/wwWVPD1ei/9RVa6Db\nvxjscE4BkcAv4ddr9szi9B11+0aLkFGgNrVSX6Oaf10PwKaZFtJFG62woSXF\nL1B+UNphFkN6ipT+8XytjhGrkKdFn2nA2nWY5CkLlb6IjokO0BgONXluVPrw\n1X+yR6Rl31ai+G7rTN4JMzh1q5cIU+fI9DX80SwOo45BqNk9uIgOUfbq5vaS\nNMsaVS3fpaB8zZscHwFEDvrZ3lyJbxVz9bgosXjSVUwYE9Bu+TnEhzvi+N3/\naBFlfACb3VY9OoYFTXnC8W5SxZjl1RfHi5lHg3+Wk5BxN0L4+fUq0nktyoVH\nsdOx1J6jzNeQfp1Q9V6WfdwPtYL3l3W+Ar8uaVoc1ZbhsPgAH/3x1p4pfwJU\ns1rTOiBwey8kCUHFbR/WeJVpOvOba6oFrKUFltNzd6zSytOb94YpbLFQAEs/\nEJ5uhxQe1nViwV6L+4LiUH9lRcWLnowSnWnoFR3FzFf0PuOdfWGAfiL0Yv/2\nxNwzgJnlhcIAUZqa7CsH2xLdEVLhkQ3uKq4CpX5Hz0SG+1rj/bsJqSojS4s0\nLtKuCK1Blpdh3YctWka9xun9WgUlmJdsvhz3G7MNjVAxmZdiD7cP57nn4rXX\nA2qfxH3m1650ErzVXijH9fqDqXMlNNIFORU19OAKHlQ5BmdpMAoGVe4YX9N8\n15F3LrOUCbeOQUvyuHFfAofo3VGHasfceyJEbORzHqtQe8bNqjx5LqI02DD8\nq0EWKy6BOkGvKujIUw54AbN+r0EP4qGjtwSv/QGUwSXoDwYyY1WAOteST+i5\nknVSuEZ1ynVIkvaeVpJsa8uX/xYDLwxIzY4wYzMwYjk2ZTMzNDI1NGY3OTFk\nZDZlMjJlYjE0Yzg0YmJmZjVjMTE0OGY5OTAxOWZmNzAwMTVlNjVmY2I0NjJk\nIDwwYzMwYjk2ZTMzNDI1NGY3OTFkZDZlMjJlYjE0Yzg0YmJmZjVjMTE0OGY5\nOTAxOWZmNzAwMTVlNjVmY2I0NjJkQGVudmtleS5jb20+wsB1BBABCAApBQJZ\ngRXnBgsJBwgDAgkQgaAUh92OkrMEFQgKAgMWAgECGQECGwMCHgEAAJqDB/9+\nL13QtKv3PGcD3zKCZYHZVrGfQf+ctEhtfsEzrCX4DavgIC4ojtE1AAPXgdBj\nUasWXJhIxuhMU/Vufa65/JwrEWPKj9aANwCHKcHTWX/4PidIF+L0oxNxGKBZ\n9SOCknks4jhEc7S0fTiC3SjLEsvs2H5PlYE3BqXvU9R11NX5r29u0RTQfeod\njeSiEz4kcCEjFT97IQFtIteU9AwP6Fwk1lIFTWjdgTYtzCb0IvtYX6RzTUBa\n1YohU6eTTCGw2x4F0/CISpp++TJl2BPPglwvy9vCbqbvWCVTH3Ryweo5e1+u\nEyxHXIm04WJaXuu8G8UcBizrXzamnRSJpDKxqUFVx8MGBFmBFecBCACmL5mX\nVuF20Tb5cxSSxwqfkYHipX3Gru5QaCMkRllWtd6RC4Y3smr2NIM6BJDtDApq\nnAWgzaesq1e5MQYTogO3gRiOtgRhIH+TpG1DXxFgx5z4CaeoL+Zyo8fLH2M2\nthI7JktJ5qbTHqUYbEqcHAz4EJk6iwXqLXRL7IN1+e4I1lD00vINzo0hhp40\nLlBN5dPTiEy9imtHdpH/Liy1IXY/oqML8+FOvnB8DmRON8M/HdnbmvuMob5Z\nkdrl+vn2eupgtLuHPegcJdYhcbTR70kJc9lTTROGjPd39m0ZRmNvUb4z/xXd\nM5n2wgz4WHJL5pxPxAVI2xlYF6HJ1R7y5PoZABEBAAH+CQMIrK7eUk3qhl1g\n6DMVpyV5SXsK6SKqHS2va3F02rGfm2lN6TIb006ZglUeoix5ipg1lJpkKBKZ\nCSCbBRy7Fj2TCoc66RFLivlW/p6oOIuu4mGQhvQ9awjgK2WxRXkEBtvt/xfk\nUilyc13G16e99sHQ1WWBeddhOdPiSnIaNbX1JU94HFCmfx+qtSQunmQF8+1l\nmbENzJPNyh6ej2Syhsm6jHiocqZwhhssouCpQWc65SblcZo9+YDc5mEF5u7V\nBoXgmmoNaRzK7qSSfHrv6nmch/a01sEt34X9MVwOGOQ4OU0UyBm5xgo4zr4L\nAPfLjtzMVl7XqXLttAYQ9d/qc3aU07L+ylDguwjFdpJhzTXGsIRFvV723wRM\nEfJUwIdDA3xBR8rANcLXyMR97Ie8Bg+U35PToYgEnFknP+4bCAj8HGIQccx8\nCDpg8/et1vHs2O12JMeDaLnIZqDA9ZpBudfd6n7v/P1hpk9fz+bO5pu1hrko\nUZZhLcrQlQQ3k5NpeFMCiddSNArVtPnEsFZ3wwrPpqf23Am+IhQc9jber7PJ\ndfUXwSPY0tKDfpcrk6jRQLUk3Kn49ylRZYY/AFQeZrrnqD8WflICDzlxfkZx\nYjEeDvg+AOVGx7bM8upDf5CeYsb/N/R7PVVGbARNc5HaABSnfzyDCt2nkRqk\nbXNkDn0CK6IAIY1BnHnFdnvxLSYhRCz8C5YNWw4p92gDU4HgRjt9ORPwnmkv\nGvJ0JSbbXsKc10tMA+UbJpdsqeFXTaNC0GvKaY+4gd/ejUsdfj3PJQvOFWTI\n79onTBF9tGHV0OQeiqRfXAIT9UbAx+NfgezzbQMkqXgzlQFSrDv0c1JNszmk\nTTP7NVas7jxvXwgD318d2APgbyhP8W8ISn7rsd3LlXuXKGU8oGoh9Qg2KU2J\nFRkk5XstJHThwsBfBBgBCAATBQJZgRXoCRCBoBSH3Y6SswIbDAAAyLcIAIhu\ntpWLS6cz8YWoROMDpKPOVVFTLr6ATLkroeXWk/aFNXEtWfIFpJbjVBcngEpe\nynaH/ensfyV5fpMca7wHkjQJtWm6d7AOrkX2dpPhTw6K/g9lFb/tndv7a/kk\nkfiyzapewHX8rZ4EKCgkES8u2fHsAwtGlQrn3fuqMGeZTwYP4SxI6Az2emvD\n3W6k0k56KU7OiCM8USrUOpq8ApEMfCtoXDSRYOMLl5i8ZLVQBu5PIORQQDLs\nsbNRox0CThvy6ad3UDu4fD/s0325Qnuk7F+uwOW0DYgeCrmZvzN8A7pq6hHU\nPTZIpmPeXld5zj6WU5E7kbBe4LlDD/rXfZAKYXY=\r\n=2eBv\r\n-----END PGP PRIVATE KEY BLOCK-----\r\n")

var rawEnvPassphrase = []byte("3e8DGyLtTuWLMxH2")
