# design-patterns-golang
Design patterns implementation using golang


https://refactoring.guru/design-patterns/catalog

Design patterns are typical solutions to commonly occurring problems in software design. 
They are like pre-made blueprints that you can customize to solve a recurring design problem in your code.

Patterns are often confused with algorithms, because both concepts describe typical solutions to some known problems. 
While an algorithm always defines a clear set of actions that can achieve some goal, a pattern is a more high-level description of a solution. 
The code of the same pattern applied to two different programs may be different.

An analogy to an algorithm is a cooking recipe: both have clear steps to achieve a goal. 
On the other hand, a pattern is more like a blueprint: you can see what the result and its features are, but the exact order of implementation is up to you.


The Catalog of Design Patterns:


1. Creational Patterns: These patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.

2. Structural Patterns: These patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.

3. Behavioral Patterns: These patterns are concerned with algorithms and the assignment of responsibilities between objects.


Creational Patterns: 

1. Factory Method: https://refactoring.guru/design-patterns/factory-method
   Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.
   Use the Factory Method when you want to provide users of your library or framework with a way to extend its internal components.
   Use the Factory Method when you want to save system resources by reusing existing objects instead of rebuilding them each time.

2. Abstract Factory: https://refactoring.guru/design-patterns/abstract-factory
    Use the Abstract Factory when your code needs to work with various families of related products, 
    but you don’t want it to depend on the concrete classes of those products—they might 
    be unknown beforehand or you simply want to allow for future extensibility.

