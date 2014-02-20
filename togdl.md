Typed OGDL (Proposal)
=====================

OGDL is a minimal textual format that represents trees or graphs. This proposal
attempt to provide a standard way to allow OGDL represent typed objects.

The design favors Go, but should also be easily implemented in other languages.

Typed OGDL has no types (almostly)
----------------------------------
The advantage of a minimal textual format over XML is that it contains no type
information but only values, making it extremely readable. This is exact why
JSON is so popular. Containing all the type information in the text would just
make the format not minimal anymore, and not better than XML.

Actually, type information has already been defined in language data structures
and can be obtained with reflections, so data structure in the language is just
like a schema, allowing type information restored given only values.

The real difficulty of this approach happens when the object is a polymorphic
one, e.g. an interface. In such cases, the concrete type can not be restored
given only the interface defined in the data structure, so we still have to
introduce a way to define concrete type optionally, e.g.

For an error object (an interface in Go):

    Err #ParseError
        Msg  "expected blablabla"
        Line 23
        File /xxx/xxx

where Err is the name of the field (in a struct), and its type is ParseError.

Fundamental Types
-----------------
Boolean, numeric types: Same as Go literals.

String, byte: OGDL node.

Zero value: nil.

Array/slice 
-----------
Array, slice is just OGDL child nodes listed one by one.

So top level is an array/slice.

There is no need to distinguish array and slice in the format.

Struct/Map
----------
The struct/map can be define starting with a minus sign "-", e.g.

    -
        FSType rootfs
        Total  5908
        Used   5112

Map is the same as struct, with keys just like field names.

Interface
---------
Store value with concrete type included.

Function/Channel
----------------
Not clear, omitted for now.
Possible approach:
Use the function's name as value, however, it means nothing if it is a function
literal (unnamed function).

Pointer
-------
If the pointer is not a circular one, it is just the same as its pointed value.

For a cyclic pointer, it should follow OGDL's way to handle it.

