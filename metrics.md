# Metrics

- Metrics always return numerical values.
- Metric can only return value based on some dimensionality/context in which it is executed
- Metrics in MAQL start with keyword SELECT.

## Create Metrics Using API

You can create and manage metrics through the workspace entity API interface.
The entity API endpoints for metrics:

- `/api/v1/entities/workspaces/{workspace-id}/metrics`
- `/api/v1/entities/workspaces/{workspace-id}/metrics/{metric-id}`

### Body Syntax for Metrics in the Entity API Interface

```json
{
  "data": {
    "attributes": {
      "title": "<metric_title>",
      "description": "<metric_description>",
      "content": {
        "format": "<number_format>",
        "maql": "<maql_expression>"
      }
    },
    "id": "<metric_id>",
    "type": "metric"
  }
}
```

### Example Request Body

```json
{
  "data": {
    "attributes": {
      "title": "Order Amount",
      "description": "Sum of order prices",
      "content": {
        "format": "$#,##0",
        "maql": "SELECT SUM({fact/order_lines.price}*{fact/order_lines.quantity})"
      }
    },
    "id": "order_amount",
    "type": "metric"
  }
}
```

## MAQL (Multidimensional Analytical Query Language)

MAQL is used to create reusable mutltidimensional queries that combine multiple facts and attributes
