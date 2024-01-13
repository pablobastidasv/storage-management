[drawer-component-url]: https://www.webcomponents.org/element/side-drawer

# Used Web Components

- [side-drawer][drawer-component-url]

# Run locally

Ensure that the file `.env` exists and have the value of the Client Secret for the authentication

```sh
AUTH_CLIENT_SECRET = "<<secret value>>"
```

# Designs

```mermaid
---
title: Storage Model Diagrama
---
classDiagram
    class InventoryItem {
        -int qty
        -InventoryProduct product
    }

    class InventoryProduct {
        -string name
        -Presentation presentation
    }

    class Product {
        -string name
        -Presentation presentation
    }

    class Presentation {
        <<enumeration>>
        QUANTITY
        KG
        GRMS
    }

    InventoryItem "1" --o "1" InventoryProduct
    InventoryProduct "*" --o "1" Presentation
    Product "*" --o "1" Presentation
    Transaction "*" --o "1" Product
    Transaction "*" --o "1" Storage
```

```mermaid
---
title: Remissions model diagram
---
classDiagram
    class Product {
    }

    class Client {
    }

    class Remission {
        -string id
        -Product product
        -int qty
        -RemissionState state
        -Client client
        -bool withReturn
        -int returnedQty
        -Time createdAt
        -Time finishedAt
    }

    class RemissionState {
        <<enumeration>>
        IN_PROGRESS
        FINISHED
    }

    Remission "*" --o "1" Client
    Remission "*" --o "1" Product
    Remission "*" --o "1" RemissionState
```
