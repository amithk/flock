# File Level Locks

Many modern day applications need "locks" that protect the access to shared resources across multiple processes. For example, multiple processes running on a same machine are trying to read, write and update a same file, and same offset within the file. In this case, the file updates from multiple processes should be consistent. 

File level locks can be used to protect against concurrent access to such shared resources/files.

This library internally makes use of "file name uniqueness" guarantees by the open/create system call. In other words, if two processes try to create a file at a same path on the file system, only one of them will succeed, and the other one will fail. 

This library provides following interfaces.

- Lock
- Unlock
- LockWithTimeout

The common input parameter for all of these APIs is 'spath', which should be a valid path for a new file to be created.

Note: The lock taken by the process is persistent (i.e. lock - if taken - persists a process crash). Application developer has a responsibility to cleanup the lock file in case of a crash.
