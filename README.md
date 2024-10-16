# Seshat - ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_R20.png?7bb17) ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_X1.png?f2a8c) ![](https://en.wikipedia.org/w/extensions/wikihiero/img/hiero_B1.png?ca40a)

# Install

- Download a version from the [releases page](https://github.com/nobe4/seshat/releases/latest).

    To find your OS and architecture see [this](https://superuser.com/a/173795).

- (optional) Rename to `seshat` and put it in your `$PATH`.

- Run

    ```sh
    chmod +x ./path/to/file
    ```

- :warning: If you are on macOS, you might need to run

    ```sh
    xattr -d com.apple.quarantine ./path/to/file
    ```

# Usage

- Write a `config.yaml` file

  See [config.yaml](./examples/config.yaml) for an example.

- Run `seshat` in the same directory as the `config.yaml` file.
- See `stdout` for the output and `output.pdf` for the generated PDF.
