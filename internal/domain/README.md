The domain layer is the layer of the clean architecture that contains the Critical Business Rules of the application. It is responsible for defining the domain entities and their relationships, as well as the business rules that govern their behavior.

In the `internal/domain` folder, you might find packages that contain the following types of code:

- `Entities`: These are the core objects of the domain, such as users, products, or orders. They are typically implemented as structs, and they contain the data and behavior that are relevant to the domain.
- `Value objects`: These are objects that represent a value in the domain, such as a currency or a date. They are typically implemented as structs, and they are used to encapsulate domain-specific logic or constraints.
- `Repositories`: These are interfaces that define the methods for storing and retrieving domain objects from a persistent storage, such as a database.
- `Services`: These are interfaces or abstract classes that define the methods for performing domain-specific operations, such as calculating a discount or generating a report.

The `internal/domain` folder might also contain other packages and files that are specific to the domain, such as utility functions or test helpers.

