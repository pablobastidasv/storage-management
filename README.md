[drawer-component-url]: https://www.webcomponents.org/element/side-drawer

# Used Web Components

- [side-drawer][drawer-component-url]

# Designs

```mermaid
---
title: Models Diagrama
---
classDiagram
    class Storage {
        -[]Item items 
    }
    class Item {
        -int qty
        -product Product
    }
    
    class Product {
        -string name
        -presentation Presentation
    }
    
    class Presentation{
        <<enumeration>>
        QUANTITY
        KG
        GRMS
    }
    
    class Transaction {
        - datetime Time
        - product Product
        - storage Storage
        - qty int
    }
    
    Storage "1" --* "*" Item
    Item "1" --o "1" Product
    Product "*" --o "1" Presentation
    
    Transaction --o Product
    Transaction --o Storage
```