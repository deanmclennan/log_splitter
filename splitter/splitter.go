package splitter

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// Check operating system

func check_os() string {
	os := os.Getenv("OS")
	return os
}

// remove duplicate strings from slice
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// Split process name from PID
func split_process_name_pid(process string) (string, string) {
	// split string into two parts
	r, err := regexp.Compile(`[\w-]+`)
	if err != nil {
		log.Fatal(err)
	}
	process_name := r.FindStringSubmatch(process)
	// fmt.Println(process_name[0])

	r, err = regexp.Compile(`\[\d+\]`)
	if err != nil {
		log.Fatal(err)
	}
	pid := r.FindStringSubmatch(process)
	// fmt.Println(pid[0])

	return process_name[0], pid[0]

}

// Open log file and find all process names

func find_all_process_names(logfile string) []string {
	// Open log file
	file, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// make slice to hold regex matches
	var matches []string

	var pids []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// print regex matches

		// fmt.Println(r.FindStringSubmatch(scanner.Text()))

		// put regex matches into slice of strings
		r, err := regexp.Compile(`[\w-]+\[\d+\]:`)
		if err != nil {
			log.Fatal(err)
		}

		matches = append(matches, r.FindStringSubmatch(scanner.Text())...)

		// remove duplicates from slice

	}

	// remove duplicates from slice
	matches = removeDuplicates(matches)

	// split process name from PID
	for i, process := range matches {
		process_name, pid := split_process_name_pid(process)
		matches[i] = process_name
		pids = append(pids, pid)

		// push process name and PID to slice
		// process_system = append(process_system, process_name)

	}

	// remove duplicates from slice
	matches = removeDuplicates(matches)

	return matches

}

// For each process in slice find match in log file and move to new file

func sort_logs_by_process(process string, logfile string) []string {

	// regex matches to slice of strings
	var matches []string

	regex_string := process + `\[\d+\]:`

	// Define the regex pattern
	pattern := regexp.MustCompile(regex_string)

	// Open the log file
	file, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Scan each line and check if it matches the regex pattern
	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			// If it does, print the whole line

			matches = append(matches, line)
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matches

}

// Check if folder exists and if not create it

func check_folder_exists(folder string) bool {
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Create folder

func create_folder(folder string) {

	err := os.Mkdir(folder, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// Make file name from process name and then write to file the lines that match the process name using sort logs by process function
func write_to_file(process string, root_folder string, logfile string) {

	folder := root_folder

	// make file name from process name

	file_name := process + ".log"

	// write to file the lines that match the process name using sort logs by process function
	log_mathces_to_write := sort_logs_by_process(process, logfile)

	// create file
	f, err := os.Create(folder + "/" + file_name)
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	for _, line := range log_mathces_to_write {
		fmt.Fprintln(f, line)
	}

	// close file
	f.Close()

}

// make root folder for process logs
func make_root_folder(root_name string) string {
	err := os.Mkdir(root_name, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return root_name
}

func Split_logs(logfile string, root_folder string) {

	// make root folder for process logs
	root_name := root_folder
	root_folder_exists := check_folder_exists(root_name)
	if root_folder_exists == false {
		make_root_folder(root_name)
	}

	// find all process names
	process_names := find_all_process_names(logfile)

	// // for each process name in slice create folder
	// for _, process := range process_names {
	// 	folder_exists := check_folder_exists(root_name + "/" + process)
	// 	if folder_exists == false {
	// 		create_folder(root_name + "/" + process)
	// 	}
	// }

	// for each process name in slice move logs to folder
	for _, process := range process_names {
		write_to_file(process+"", root_name, logfile)
	}

}
