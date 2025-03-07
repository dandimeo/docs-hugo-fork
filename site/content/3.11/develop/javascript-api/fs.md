---
title: Filesystem Module
menuTitle: fs
weight: 45
description: >-
  Filesystem Module
archetype: default
---
`const fs = require('fs')`

The implementation tries to follow the CommonJS
[Filesystem/A/0](http://wiki.commonjs.org/wiki/Filesystem/A/0)
specification where possible.

## Working Directory
The directory functions below shouldn't use the current working directory of the server like `.` or `./test`.
You will not be able to tell whether the environment the server is running in will permit directory listing,
reading or writing of files.

You should either base your directories with `getTempPath()`, or as a Foxx service use the
[module.context.basePath](../foxx-microservices/reference/service-context.md).

## Single File Directory Manipulation

### exists

checks if a file of any type or directory exists
`fs.exists(path)`

Returns true if a file (of any type) or a directory exists at a given
path. If the file is a broken symbolic link, returns false.

### isFile

tests if path is a file
`fs.isFile(path)`

Returns true if the *path* points to a file.

### isDirectory

tests if path is a directory
`fs.isDirectory(path)`

Returns true if the *path* points to a directory.

### size

gets the size of a file
`fs.size(path)`

Returns the size of the file specified by *path*.

### mtime

gets the last modification time of a file
`fs.mtime(filename)`

Returns the last modification date of the specified file. The date is
returned as a Unix timestamp (number of seconds elapsed since January 1 1970).

### pathSeparator
`fs.pathSeparator`

If you want to combine two paths you can use fs.pathSeparator instead of `/` or `\\`.

### join
`fs.join(path, filename)`

The function returns the combination of the path and filename, e.g.
`fs.join('folder', 'file.ext')` would return `folder/file.ext`.

### getTempFile

returns the name for a (new) temporary file
`fs.getTempFile(directory, createFile)`

Returns the name for a new temporary file in directory *directory*.
If *createFile* is *true*, an empty file will be created so no other
process can create a file of the same name.

**Note**: The directory *directory* must exist.

### getTempPath

returns the temporary directory
`fs.getTempPath()`

Returns the absolute path of the temporary directory

### makeAbsolute

makes a given path absolute
`fs.makeAbsolute(path)`

Returns the given string if it is an absolute path, otherwise an
absolute path to the same location is returned.

### chmod

sets file permissions of specified files (non windows only)
`fs.chmod(path, mode)`

where `mode` is a string with a leading zero matching the `OCTAL-MODE` as explained
in *nix `man chmod`.

Returns true on success.

### list

returns the directory listing
`fs.list(path)`

The functions returns the names of all the files in a directory, in
lexically sorted order. Throws an exception if the directory cannot be
traversed (or path is not a directory).

**Note**: this means that list("x") of a directory containing "a" and "b" would
return ["a", "b"], not ["x/a", "x/b"].

### listTree

returns the directory tree
`fs.listTree(path)`

The function returns an array that starts with the given path, and all of
the paths relative to the given path, discovered by a depth first traversal
of every directory in any visited directory, reporting but not traversing
symbolic links to directories. The first path is always `""`, the path
relative to itself.

### makeDirectory

creates a directory
`fs.makeDirectory(path)`

Creates the directory specified by *path*.

### makeDirectoryRecursive

creates a directory
`fs.makeDirectoryRecursive(path)`

Creates the directory hierarchy specified by *path*.

### remove

removes a file
`fs.remove(filename)`

Removes the file *filename* at the given path. Throws an exception if the
path corresponds to anything that is not a file or a symbolic link. If
"path" refers to a symbolic link, removes the symbolic link.

### removeDirectory

removes an empty directory
`fs.removeDirectory(path)`

Removes a directory if it is empty. Throws an exception if the path is not
an empty directory.

### removeDirectoryRecursive

removes a directory
`fs.removeDirectoryRecursive(path)`

Removes a directory with all subelements. Throws an exception if the path
is not a directory.

## File IO

### read

reads in a file
`fs.read(filename)`

Reads in a file and returns the content as string. Please note that the
file content must be encoded in UTF-8.

### read64

reads in a file as base64
`fs.read64(filename)`

Reads in a file and returns the content as string. The file content is
Base64 encoded.

### readBuffer

reads in a file
`fs.readBuffer(filename)`

Reads in a file and returns its content in a Buffer object.

### readFileSync

`fs.readFileSync(filename, encoding)`

Reads the contents of the file specified in `filename`. If `encoding` is specified,
the file contents will be returned as a string. Supported encodings are:
- `utf8` or `utf-8`
- `ascii`
- `base64`
- `ucs2` or `ucs-2`
- `utf16le` or `utf16be`
- `hex`

If no `encoding` is specified, the file contents will be returned in a Buffer
object.

### write
`fs.write(filename, content)`

Writes the content into a file. Content can be a string or a Buffer
object. If the file already exists, it is truncated.

### writeFileSync

`fs.writeFileSync(filename, content)`

This is an alias for `fs.write(filename, content)`.

### append

`fs.append(filename, content)`

Writes the content into a file. Content can be a string or a Buffer
object. If the file already exists, the content is appended at the
end.

## Recursive Manipulation

### copyRecursive

copies a directory structure
`fs.copyRecursive(source, destination)`

Copies *source* to *destination*.
Exceptions will be thrown on:
 - Failure to copy the file
 - specifying a directory for destination when source is a file
 - specifying a directory as source and destination

### CopyFile

copies a file into a target file
`fs.copyFile(source, destination)`

Copies *source* to destination. If Destination is a directory, a file
of the same name will be created in that directory, else the copy will get
the specified filename.

### linkFile

creates a symbolic link from a target in the place of linkpath.
`fs.linkFile(target, linkpath)`

In `linkpath` a symbolic link to `target` will be created.

### move

renames a file
`fs.move(source, destination)`

Moves *source* to destination. Failure to move the file, or
specifying a directory for destination when source is a file will throw an
exception. Likewise, specifying a directory as source and destination will
fail.

## ZIP

### unzipFile

unzips a file
`fs.unzipFile(filename, outpath, skipPaths, overwrite, password)`

Unzips the zip file specified by *filename* into the path specified by
*outpath*. Overwrites any existing target files if *overwrite* is set
to *true*.

Returns *true* if the file was unzipped successfully.

### zipFile

zips a file
`fs.zipFile(filename, chdir, files, password)`

Stores the files specified by *files* in the zip file *filename*. If
the file *filename* already exists, an error is thrown. The list of input
files *files* must be given as a list of absolute filenames. If *chdir* is
not empty, the *chdir* prefix will be stripped from the filename in the
zip file, so when it is unzipped filenames will be relative.
Specifying a password is optional.

Returns *true* if the file was zipped successfully.
