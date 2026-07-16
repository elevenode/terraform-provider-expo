resource "expo_update_branch" "production" {
  app_id = expo_app.this.id
  name   = "production"
}

resource "expo_update_branch" "next" {
  app_id = expo_app.this.id
  name   = "next"
}

# Routes every update on the channel to a single branch.
resource "expo_update_channel" "production" {
  app_id = expo_app.this.id
  name   = "production"

  branch_mapping = jsonencode({
    version = 0
    data = [{
      branchId           = expo_update_branch.production.id
      branchMappingLogic = "true"
    }]
  })
}

# Rolls 10% of traffic onto the "next" branch, falling back to production.
resource "expo_update_channel" "rollout" {
  app_id = expo_app.this.id
  name   = "rollout"

  branch_mapping = jsonencode({
    version = 0
    data = [
      {
        branchId = expo_update_branch.next.id
        branchMappingLogic = {
          operand               = 0.1
          clientKey             = "rolloutToken"
          branchMappingOperator = "hash_lt"
        }
      },
      {
        branchId           = expo_update_branch.production.id
        branchMappingLogic = "true"
      },
    ]
  })
}
