# errors [![GoDoc](https://godoc.org/github.com/boreq/errors?status.svg)](https://godoc.org/github.com/boreq/errors)

Because there is a limit to this insanity.

## What is this?

This library extends the `errors` package from the standard library with two
well-known methods: `Wrap` and `Wrapf`.

    func Wrap(err error, message string) error {
    	...
    }
    
    func Wrapf(err error, format string, args ...interface{}) error {
    	...
    }

## Why?

The Go Blog helpfully recommends writing the following in order to wrap an
error:

    if err != nil {
        return fmt.Errorf("decompress %v: %w", name, err)
    }

Well, my (and hopefully your) response to this is "*No, thank you Mr. Pike!*".

Thanks to this library you can continue using the old approach:

    if err != nil {
        return errors.Wrapf(err, "decompress %v", name)
    }

The second version is in my opinion much more readable as well as idiomatic -
and as far as I know the only reason for not going with this approach in the
standard library was the aversion to forcing every single Go program in the
world to import `fmt` whenever `errors` is imported. There are other funny
situations in the standard library caused by avoiding dependencies.

## Why not `github.com/pkg/errors`?

Unfortunately `github.com/pkg/errors` is incompatibile with the new methods
available in the standard library such as `errors.Is(...)`.
