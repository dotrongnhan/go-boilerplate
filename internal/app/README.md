Not all business rules are as pure as `Entities`. Some business rules make or save money for the business by defining and constraining the way that an automated system operates. These rules would not be used in a manual environment, because they make sense only as part of an automated system.

The application layer is the layer of the clean architecture that sits between the domain layer and the infrastructure layer. It is responsible for coordinating the activities of the other layers, and for implementing any application-specific business logic that is not part of the domain layer.

In the `internal/app` folder, you might find packages that contain the following types of code:

- `Use cases`: These are the specific actions that the application can perform, such as creating a new user or updating a product. Each use case is typically implemented as a function or method that takes input, performs some processing, and returns output.
- `Services`: These are classes or structs that implement the use cases, by calling the appropriate domain and infrastructure functions and methods.
- `Interactors`: These are classes or structs that handle the communication between the use cases and the domain layer. They are responsible for transforming the input and output of the use cases into a form that can be understood by the domain layer, and vice versa.
- `Request and response models`: These are structs that define the input and output of the use cases. They are typically used to validate the input, and to structure the output in a way that is easy to understand and use.

The `internal/app` folder might also contain other packages and files that are specific to the application, such as utility functions or test helpers.