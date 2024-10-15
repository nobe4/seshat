# Seshat - ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_R20.png?7bb17) ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_X1.png?f2a8c) ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_B1.png?ca40a)

# Install

- Download a version from the [releases
page](https://github.com/nobe4/seshat/releases/latest).
- Rename to `seshat` and put it in your `$PATH`.


- :warning: If you are on macOS, you might need to run

    ```sh
    xattr -d com.apple.quarantine ./seshat
    ```


# Usage

- Write a `rules.yaml` file

  See [rules.yaml](./rules.yaml) for an example.

- Run `seshat` in the same directory as the `rules.yaml` file.
- See `stdout` for the output and `test.pdf` for the generated PDF.
