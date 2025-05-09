# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    local = {
      source = "hashicorp/local"
    }
    equinix = {
      source = "equinix/equinix"
    }
    time = {
      source = "hashicorp/time"
    }
    random = {
      source = "hashicorp/random"
    }
    tls = {
      source = "hashicorp/tls"
    }
  }
  required_version = ">= 0.13"
}
