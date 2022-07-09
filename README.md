# xdg-direct

`xdg-direct` was written to deal with the situation where you click on a link, _and it opens the browser in the window with the wrong profile_. And no, I don't want to click on which profile to open a link in each time.

The way it deals with this is by acting as a glob-based "router" between `xdg-open` and the browser. It scans the opened url for patterns defined in the config (`/etc/xdg-direct.config`) and runs the command associated with the url. See the `Usage` section for examples.

## Usage

Let's say you have a Firefox profile called `work` and the following config file:

```yaml
# /etc/xdg-direct/config.yaml
default_command: "firefox"
mappings:
  - url: "*.gitlab.org/*"
    cmd: "firefox"
    args:
      - "-P"
      - "work"
debug_mode: false
```

When running `xdg-direct "https://example.gitlab.org/example"`, a match with `*.gitlab.org/*` will be found, and the command `firefox -P work` will be run.
The url will automatically be included at the _end_ of the command. 
When there is no match with any rule, say with `https://wikipedia.org`, the default command will be run (just `firefox`in this case).

### Use with `xdg-open`

The Makefile provides a convenient way to install both the programme and the `.desktop` application file needed for running it with `xdg-open`; 
this can be done by running `make alterXdg`. When running without `make`, run the scripts from the Makefile in order in this repo's root directory.


## TODO

- add automatic build
- add some unit tests
- verify this works for MacOS
