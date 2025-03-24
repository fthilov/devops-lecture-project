terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "4.24.0"
    }
  }
}

provider "azurerm" {
    subscription_id = "0490a39a-c7e6-4bd7-8cbe-abb171cd37e1"
    features {}
}


resource "azurerm_resource_group"  "tofu-ressource" {
  name                = "tofu-ressource"
  location            = var.region
}

resource "azurerm_kubernetes_cluster" "kube-cluster" {
  name                = "kube-cluster"
  location            = var.region
  resource_group_name = "tofu-ressource"
  dns_prefix          = "test1"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }
}