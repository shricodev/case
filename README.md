## Plan

```bash
dircase lower ./my-folder
dircase upper ./my-folder
dircase title ./my-folder
dircase camel ./my-folder

dircase lower ./my-folder --recursive
dircase lower ./my-folder --target dirs
dircase lower ./my-folder --target files
dircase lower ./my-folder --target all
dircase lower ./my-folder --dry-run
```

For the first version, I’d keep it focused:

```bash
dircase lower <path>
dircase upper <path>
dircase capitalize <path>
dircase title <path>
```

Use flags like:

```bash
-r, --recursive
-n, --dry-run
--target dirs|files|all
--include-hidden
--preserve-ext
```

For Go tooling, I’d use:

```txt
cobra      -> CLI commands and flags
filepath   -> path handling
os.ReadDir -> reading directories
os.Rename  -> renaming
strings    -> lower/upper/title-ish transforms
```

Cobra is a solid choice for this kind of CLI because it is made for Git-style commands like `dircase lower` and `dircase upper`, and it is used by tools like Kubernetes, Hugo, and GitHub CLI. ([GitHub][1])

Suggested project structure:

```txt
dircase/
  cmd/
    root.go
    lower.go
    upper.go
    capitalize.go
  internal/
    caseconv/
      caseconv.go
    walker/
      walker.go
    rename/
      rename.go
  main.go
  go.mod
```

I’d separate the logic like this:

```txt
cmd/         -> CLI stuff only
caseconv/    -> lowercase, uppercase, capitalize, title
walker/      -> recursively find dirs/files
rename/      -> collision checks, dry-run, os.Rename
```

The important thing: don’t mix traversal, renaming, and case conversion in one big file. Future you will hate that.

For recursive directory renaming, be careful. If you rename folders while walking top-down, paths can break because the parent directory name changes before children are processed. A good pattern is:

```txt
1. collect all matching paths first
2. sort deepest paths first
3. rename from bottom to top
```

Go has `filepath.WalkDir` for recursively walking a tree, but for renaming directories, collecting first and then renaming deepest-first is safer. `WalkDir` walks the file tree from a root and gives you each file or directory path. ([Go Packages][2])

Core types could look like this:

```go
type CaseMode string

const (
    Lower      CaseMode = "lower"
    Upper      CaseMode = "upper"
    Capitalize CaseMode = "capitalize"
    Title      CaseMode = "title"
)

type Target string

const (
    TargetDirs  Target = "dirs"
    TargetFiles Target = "files"
    TargetAll   Target = "all"
)

type Options struct {
    Root          string
    Mode          CaseMode
    Target        Target
    Recursive     bool
    DryRun        bool
    IncludeHidden bool
    PreserveExt   bool
}
```

I’d make the first release do this:

```bash
dircase lower . --recursive --dry-run
dircase lower . --recursive
dircase upper ./packages --target dirs
dircase capitalize ./content --target dirs
```

And later you can add aliases:

```bash
dircase l .
dircase u .
dircase cap .
```

Tiny UX detail that would make it feel polished:

```bash
dircase lower . -r --dry-run
```

Output:

```txt
DRY RUN

./My Folder       -> ./my folder
./Blog Posts      -> ./blog posts
./API Docs        -> ./api docs

3 directories would be renamed.
```

My personal recommendation:

Use **Cobra**, name the binary **`dircase`**, start with commands like **`lower`**, **`upper`**, **`capitalize`**, and keep `--target files|dirs|all` ready from day one. That way it starts as a directory tool but doesn’t paint you into a corner.

[1]: https://github.com/spf13/cobra?utm_source=chatgpt.com "spf13/cobra: A Commander for modern Go CLI interactions"
[2]: https://pkg.go.dev/path/filepath?utm_source=chatgpt.com "path/filepath"
