resource "null_resource" "example" {
  provisioner "local-exec" {
    command = "echo 'This is another echo command.'"
  }
}
