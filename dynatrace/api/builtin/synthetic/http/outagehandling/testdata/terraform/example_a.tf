resource "dynatrace_http_monitor_outage" "#name#" {
  global_consecutive_outage_count_threshold = 3
  global_outages                            = true
  local_consecutive_outage_count_threshold  = 5
  local_location_outage_count_threshold     = 1
  local_outages                             = true
  scope                                     = "environment"
}
