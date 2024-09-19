variable "client_id" {
  description = "The Client ID of the Azure Service Principal"
}

variable "client_secret" {
  description = "The Client Secret of the Azure Service Principal"
}

variable "subscription_id" {
  description = "The Subscription ID in Azure"
}

variable "tenant_id" {
  description = "The Tenant ID in Azure"
}

variable "resource_group_name" {
  description = "The name of the resource group"
  default     = "azr-rg-challenge-n"
}

variable "location" {
  description = "The Azure region to create resources in"
  default     = "East US"
}

variable "aks_cluster_name" {
  description = "The name of the AKS cluster"
  default     = "azr-aks-challenge-n"
}

variable "node_count" {
  description = "The number of nodes in the AKS cluster"
  default     = 3
}

variable "node_size" {
  description = "The size of the nodes in the AKS cluster"
  default     = "Standard_D2as_v5"
}