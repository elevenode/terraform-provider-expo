resource "expo_android_keystore" "release" {
  keystore_base64   = filebase64("${path.module}/release.jks")
  keystore_password = var.keystore_password
  key_alias         = "upload"
  key_password      = var.key_password
}

output "release_sha1" {
  value = expo_android_keystore.release.sha1_certificate_fingerprint
}
