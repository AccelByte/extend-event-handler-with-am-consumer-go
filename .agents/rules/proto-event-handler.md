# Proto Files (Event Handler)

Event proto files are provided by AccelByte and define the event message schemas.

- To handle a new event type, add the corresponding AccelByte AsyncAPI proto file and implement a new service handler.
- Do **not** modify existing event proto definitions — they must match what AGS publishes.
- Generated code is regenerated from proto and should not be hand-edited.
