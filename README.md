# genpass

**genpass** is a small utility to help generate random passwords, it has no external dependencies. It's possible to choose the length of the password and there are 4 available alphabets to choose from (letters, letters+digits, letters+digits+punctuation and digits).



## Installation

    $ go get -u github.com/mjvm/genpass

## Usage

    $ genpass -h
    Usage of genpass:
        -a uint
            alphabet used to generate the password, options ASCII letters only: 0, ASCII letters and numbers: 1, ASCII letters, numbers and punctuation: 2, numbers only: 3. (default 1)
        -n
            uint length of the password (default 32)

Running `genpass` without options will generate `10` passwords based on the default alphabet (letters+digits)

    $ genpass
    p7RmQ8CwWbNXvvF8MZh2u9pVN4rYWo3e
    ...
    48zbWQi9M9sjY8bDYTE7NPBO8PXBxnUZ

Changing password's length and the alphabet

    $ genpass -n 10 -a 2
    UY6b%
    ...
    /?&X#
