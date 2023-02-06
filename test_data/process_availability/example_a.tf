resource "dynatrace_process_availability" "example" {
  enabled = true
  name    = "example"
  rules {
    rule {
      property  = "executable"
      condition = "$contains(svc)"
    }
  }
  metadata {
    item {
      key   = "foo"
      value = "bar"
    }
  }
}