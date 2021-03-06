// Code generated by protoc-gen-goext. DO NOT EDIT.

package kms

func (m *SymmetricEncryptRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricEncryptRequest) SetVersionId(v string) {
	m.VersionId = v
}

func (m *SymmetricEncryptRequest) SetAadContext(v []byte) {
	m.AadContext = v
}

func (m *SymmetricEncryptRequest) SetPlaintext(v []byte) {
	m.Plaintext = v
}

func (m *SymmetricEncryptResponse) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricEncryptResponse) SetVersionId(v string) {
	m.VersionId = v
}

func (m *SymmetricEncryptResponse) SetCiphertext(v []byte) {
	m.Ciphertext = v
}

func (m *SymmetricDecryptRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricDecryptRequest) SetAadContext(v []byte) {
	m.AadContext = v
}

func (m *SymmetricDecryptRequest) SetCiphertext(v []byte) {
	m.Ciphertext = v
}

func (m *SymmetricDecryptResponse) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricDecryptResponse) SetVersionId(v string) {
	m.VersionId = v
}

func (m *SymmetricDecryptResponse) SetPlaintext(v []byte) {
	m.Plaintext = v
}

func (m *GenerateDataKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *GenerateDataKeyRequest) SetVersionId(v string) {
	m.VersionId = v
}

func (m *GenerateDataKeyRequest) SetAadContext(v []byte) {
	m.AadContext = v
}

func (m *GenerateDataKeyRequest) SetDataKeySpec(v SymmetricAlgorithm) {
	m.DataKeySpec = v
}

func (m *GenerateDataKeyRequest) SetSkipPlaintext(v bool) {
	m.SkipPlaintext = v
}

func (m *GenerateDataKeyResponse) SetKeyId(v string) {
	m.KeyId = v
}

func (m *GenerateDataKeyResponse) SetVersionId(v string) {
	m.VersionId = v
}

func (m *GenerateDataKeyResponse) SetDataKeyPlaintext(v []byte) {
	m.DataKeyPlaintext = v
}

func (m *GenerateDataKeyResponse) SetDataKeyCiphertext(v []byte) {
	m.DataKeyCiphertext = v
}

func (m *SymmetricReEncryptRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricReEncryptRequest) SetVersionId(v string) {
	m.VersionId = v
}

func (m *SymmetricReEncryptRequest) SetAadContext(v []byte) {
	m.AadContext = v
}

func (m *SymmetricReEncryptRequest) SetSourceKeyId(v string) {
	m.SourceKeyId = v
}

func (m *SymmetricReEncryptRequest) SetSourceAadContext(v []byte) {
	m.SourceAadContext = v
}

func (m *SymmetricReEncryptRequest) SetCiphertext(v []byte) {
	m.Ciphertext = v
}

func (m *SymmetricReEncryptResponse) SetKeyId(v string) {
	m.KeyId = v
}

func (m *SymmetricReEncryptResponse) SetVersionId(v string) {
	m.VersionId = v
}

func (m *SymmetricReEncryptResponse) SetSourceKeyId(v string) {
	m.SourceKeyId = v
}

func (m *SymmetricReEncryptResponse) SetSourceVersionId(v string) {
	m.SourceVersionId = v
}

func (m *SymmetricReEncryptResponse) SetCiphertext(v []byte) {
	m.Ciphertext = v
}
