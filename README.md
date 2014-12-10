aliyun-sdk-go
=============

Aliyun services SDK in golang.

This SDK is aim to implement all the APIs of Aliyun.

Notice
------
Most of the methods of ECS client are not tested yet (Because of unsafe operations and bills).

BE CAREFUL WITH THEM.

Note for ECS API
----------------
Because of the mistakes in Aliyun API Doc,
 some of the struct for API result have some additional wrapping.

*_Array: There should be a single node which contains only one array-type field.

*_Type:  Type name is already used by a field.
