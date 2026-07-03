package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	register(subcommand{
		name:        "compare",
		description: "Compare two .env files",
		run:         runDiff,
	})
}

func runDiff(args []string) {
	fs := newFlags("diff", args)

	fs.Usage = func() {
		fmt.Println("Usage: envdiff compare <file1> <file2>")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  envdiff compare .env.example .env")
	}

	fs.Parse(args)

	if len(fs.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "Error: two files required")
		fmt.Println()
		fs.Usage()
		os.Exit(1)
	}

	file1 := fs.Args()[0]
	file2 := fs.Args()[1]

	keys1, err := parseEnvFile(file1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file1, err)
		os.Exit(1)
	}

	keys2, err := parseEnvFile(file2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file2, err)
		os.Exit(1)
	}

	printDiff(file1, file2, keys1, keys2)
}

func parseEnvFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	keys := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		keys[key] = value
	}

	return keys, scanner.Err()
}

func printDiff(file1, file2 string, keys1, keys2 map[string]string) {
	missing := []string{}
	extra := []string{}
	changed := []string{}

	// keys in file1 but missing in file2
	for key := range keys1 {
		if _, exists := keys2[key]; !exists {
			missing = append(missing, key)
		}
	}

	// keys in file2 but not in file1
	for key := range keys2 {
		if _, exists := keys1[key]; !exists {
			extra = append(extra, key)
		}
	}

	// keys in both but different values
	for key, val1 := range keys1 {
		if val2, exists := keys2[key]; exists && val1 != val2 {
			changed = append(changed, key)
		}
	}

	if len(missing) == 0 && len(extra) == 0 && len(changed) == 0 {
		fmt.Println("✓ No differences found")
		return
	}

	fmt.Printf("Comparing %s → %s\n\n", file1, file2)

	for _, key := range missing {
		fmt.Printf("  - %-30s (in %s, missing in %s)\n", key, file1, file2)
	}

	for _, key := range extra {
		fmt.Printf("  + %-30s (in %s, not in %s)\n", key, file2, file1)
	}

	for _, key := range changed {
		fmt.Printf("  ~ %-30s (different value)\n", key)
	}
}