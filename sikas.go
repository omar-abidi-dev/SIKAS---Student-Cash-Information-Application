package main

import "fmt"

// SIKAS - Student Cash Information Application
// Developed by: Omar Abidi - Gama Hendra Guntara
// Class: IF-49-INT

// ============================================================
// CONSTANTS & TYPE DEFINITIONS
// ============================================================

const NMAX int = 100 // maximum number of students

type Student struct { // composite type for one student
	nim         string // student ID
	name        string // student full name
	totalDues   int    // total dues owed in IDR
	paidAmount  int    // total amount paid so far
	lastPayDate string // date of most recent payment e.g. "2026-05-31"
	hasPaid     bool   // true if current month dues fully paid
}

type StudentArray [NMAX]Student // fixed-size array of students

// ============================================================
// INPUT / OUTPUT HELPERS
// ============================================================

// nimExists returns true if the NIM is already in the array
func nimExists(T StudentArray, n int, nim string) bool {
	found := false
	i := 0
	for i < n && !found {
		found = T[i].nim == nim
		i = i + 1
	}
	return found
}

// readStudent reads one student's data from stdin
func readStudent() Student {
	var s Student
	fmt.Print("  NIM        : ")
	fmt.Scan(&s.nim)
	fmt.Print("  Name       : ")
	fmt.Scan(&s.name)
	fmt.Print("  Dues (IDR) : ")
	fmt.Scan(&s.totalDues)
	fmt.Print("  Paid (IDR) : ")
	fmt.Scan(&s.paidAmount)
	fmt.Print("  Pay Date   : ")
	fmt.Scan(&s.lastPayDate)
	s.hasPaid = s.paidAmount >= s.totalDues
	return s
}

// printHeader prints the application banner
func printHeader() {
	fmt.Println("+++ SIKAS - Student Cash Information Application +++")
}

// printStudent displays one student record
func printStudent(s Student) {
	status := "❌ UNPAID"
	if s.hasPaid {
		status = "✅ PAID  "
	}
	fmt.Printf("  %-12s | %-20s | Dues: %8d | Paid: %8d | %-9s | %s\n",
		s.nim, s.name, s.totalDues, s.paidAmount, status, s.lastPayDate)
}

// printAllStudents displays every student in the array
func printAllStudents(T StudentArray, n int) {
	i := 0
	for i < n {
		printStudent(T[i])
		i = i + 1
	}
}

// ============================================================
// CRUD OPERATIONS
// ============================================================

// addStudent appends a new student to the array
// addStudent appends new students to the array
func addStudent(T *StudentArray, n *int) {
	var num int
	fmt.Print("Number of students to add: ")
	fmt.Scan(&num)
	i := 1
	for i <= num {
		if *n < NMAX {
			fmt.Println("--- Add Student ---")
			var s Student
			valid := false
			for !valid { // keep asking until a unique NIM is entered
				s = readStudent()
				if s.nim == "" {
					fmt.Println("NIM cannot be empty. Try again.")
				} else if nimExists(*T, *n, s.nim) {
					fmt.Println("NIM already exists. Please enter a unique NIM.")
				} else {
					valid = true
				}
			}
			T[*n] = s // add to array at index n
			*n = *n + 1
			fmt.Println("Student added successfully.")
		} else {
			fmt.Println("Storage full. Cannot add more students.")
			i = num
		}
		i = i + 1
	}
}

// seqSearchByNIM searches linearly for a student by NIM; returns index or -1
// Sequential Search
func seqSearchByNIM(T StudentArray, n int, nim string) int {
	found := -1
	i := 0
	for i < n && found == -1 {
		if T[i].nim == nim {
			found = i
		}
		i = i + 1
	}
	return found
}

// editStudent changes the data of an existing student identified by NIM
func editStudent(T *StudentArray, n int) {
	fmt.Println("--- Edit Student ---")
	fmt.Print("Enter NIM to edit: ")
	var nim string
	fmt.Scan(&nim)
	idx := seqSearchByNIM(*T, n, nim)
	if idx == -1 {
		fmt.Println("Student not found.")
	} else {
		fmt.Println("Current data:")
		printStudent(T[idx])
		fmt.Println("Enter new data:")
		T[idx] = readStudent()
		fmt.Println("Student data updated.")
	}
}

// deleteStudent removes a student by shifting the array left
func deleteStudent(T *StudentArray, n *int) {
	fmt.Println("--- Delete Student ---")
	fmt.Print("Enter NIM to delete: ")
	var nim string
	fmt.Scan(&nim)
	idx := seqSearchByNIM(*T, *n, nim)
	if idx == -1 {
		fmt.Println("Student not found.")
	} else {
		// shift elements left starting from idx
		j := idx
		for j < *n-1 { // while-loop shift
			T[j] = T[j+1]
			j = j + 1
		}
		*n = *n - 1
		fmt.Println("Student deleted successfully.")
	}
}

// recordPayment updates the payment record for a student
func recordPayment(T *StudentArray, n int) {
	fmt.Println("--- Record Payment ---")
	fmt.Print("Enter NIM: ")
	var nim string
	fmt.Scan(&nim)
	idx := seqSearchByNIM(*T, n, nim)
	if idx == -1 {
		fmt.Println("Student not found.")
	} else {
		var amount int
		var date string
		fmt.Print("Payment amount (IDR): ")
		fmt.Scan(&amount)
		fmt.Print("Payment date        : ")
		fmt.Scan(&date)
		T[idx].paidAmount = T[idx].paidAmount + amount
		T[idx].lastPayDate = date
		T[idx].hasPaid = T[idx].paidAmount >= T[idx].totalDues
		fmt.Println("Payment recorded.")
		printStudent(T[idx])
	}
}

// ============================================================
// SEARCHING
// ============================================================

// seqSearchUnpaid finds all students who have not fully paid; prints them
// Sequential Search on structured data
func seqSearchUnpaid(T StudentArray, n int) {
	fmt.Println("--- Sequential Search: Unpaid Students ---")
	found := false
	i := 0
	for i < n { // scan entire array — sequential
		if !T[i].hasPaid {
			printStudent(T[i])
			found = true
		}
		i = i + 1
	}
	if !found {
		fmt.Println("All students have paid their dues.")
	}
}

// sortByNIMAscending sorts T in-place ascending by NIM using Selection Sort
// needed to enable binary search on NIM
func sortByNIMAscending(T *StudentArray, n int) {
	i := 1
	for i <= n-1 { // outer while-loop — Selection Sort
		idxMin := i - 1
		j := i
		for j < n {
			if T[idxMin].nim > T[j].nim {
				idxMin = j
			}
			j = j + 1
		}
		t := T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

// binarySearchUnpaid uses Binary Search on a NIM-sorted array to locate a student,
// then checks payment status
func binarySearchUnpaid(T StudentArray, n int) {
	fmt.Println("--- Binary Search: Check Payment Status by NIM ---")
	fmt.Println("(Array will be sorted by NIM before searching)")

	// sort a local copy so the original order is preserved for display
	var sorted StudentArray
	i := 0
	for i < n { // copy loop
		sorted[i] = T[i]
		i = i + 1
	}
	sortByNIMAscending(&sorted, n)

	fmt.Print("Enter NIM to search: ")
	var nim string
	fmt.Scan(&nim)

	// Binary Search — ascending NIM
	kr := 0
	kn := n - 1
	found := -1
	for kr <= kn && found == -1 { // one-exit while-loop
		med := (kr + kn) / 2
		if nim < sorted[med].nim {
			kn = med - 1
		} else if nim > sorted[med].nim {
			kr = med + 1
		} else {
			found = med
		}
	}

	if found == -1 {
		fmt.Println("Student not found.")
	} else {
		printStudent(sorted[found])
		if !sorted[found].hasPaid {
			fmt.Println("Status: UNPAID — dues not fully settled.")
		} else {
			fmt.Println("Status: PAID — dues fully settled.")
		}
	}
}

// ============================================================
// SORTING
// ============================================================

// selectionSortByName sorts T ascending by student name
// Selection Sort
func selectionSortByName(T *StudentArray, n int) {
	i := 1
	for i <= n-1 { // outer while-loop (omar abidi)
		idxMin := i - 1
		j := i
		for j < n {
			if T[idxMin].name > T[j].name {
				idxMin = j
			}
			j = j + 1
		}
		t := T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

// insertionSortByArrears sorts T descending by arrears (dues − paid)
// Insertion Sort
func insertionSortByArrears(T *StudentArray, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		arrTemp := temp.totalDues - temp.paidAmount // arrears of element being inserted
		for j > 0 && arrTemp > (T[j-1].totalDues-T[j-1].paidAmount) {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

// sortMenu handles the sorting sub-menu
func sortMenu(T *StudentArray, n int) {
	fmt.Println("--- Sort Students ---")
	fmt.Println("1. Sort by name (Selection Sort, ascending)")
	fmt.Println("2. Sort by total arrears (Insertion Sort, descending)")
	fmt.Print("Choice: ")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		selectionSortByName(T, n)
		fmt.Println("Sorted by name (ascending):")
		printAllStudents(*T, n)
	} else if choice == 2 {
		insertionSortByArrears(T, n)
		fmt.Println("Sorted by arrears (highest first):")
		printAllStudents(*T, n)
	} else {
		fmt.Println("Invalid choice.")
	}
}

// ============================================================
// STATISTICS
// ============================================================

// showStatistics displays total cash collected and number of paying students
func showStatistics(T StudentArray, n int) {
	printHeader()
	fmt.Println("--- Statistics ---")
	totalBalance := 0
	paidCount := 0
	i := 0
	for i < n {
		totalBalance = totalBalance + T[i].paidAmount
		if T[i].hasPaid {
			paidCount = paidCount + 1
		}
		i = i + 1
	}
	fmt.Printf("Total students       : %d\n", n)
	fmt.Printf("Students who paid    : %d\n", paidCount)
	fmt.Printf("Students who unpaid  : %d\n", n-paidCount)
	fmt.Printf("Total cash collected : IDR %d\n", totalBalance)
}

// ============================================================
// MAIN MENU
// ============================================================

// printMenu displays the main menu
func printMenu() {
	fmt.Println("+++ SIKAS - Student Cash Information Application +++")
	fmt.Println("┌─────────────────────────────────────┐")
	fmt.Println("│  1. ➕  Add student                  │")
	fmt.Println("│  2. ✏️   Edit student                │")
	fmt.Println("│  3. 🗑️   Delete student              │")
	fmt.Println("│  4. 💰  Record payment               │")
	fmt.Println("│  5. 📋  Show all students            │")
	fmt.Println("│  6. 🔍  Search unpaid (Sequential)   │")
	fmt.Println("│  7. 🎯  Search by NIM (Binary)       │")
	fmt.Println("│  8. 🔃  Sort students                │")
	fmt.Println("│  9. 📊  Show statistics              │")
	fmt.Println("│  0. 🚪  Exit                         │")
	fmt.Println("└─────────────────────────────────────┘")
	fmt.Print("👉 Choice: ")
}

// =====================DECORATIVE FUNCTIONS====================
// =============================================================
func printSplash() { // splash screen with ASCII art
	fmt.Println()
	fmt.Println("  ███████╗██╗██╗  ██╗ █████╗ ███████╗")
	fmt.Println("  ██╔════╝██║██║ ██╔╝██╔══██╗██╔════╝")
	fmt.Println("  ███████╗██║█████╔╝ ███████║███████╗")
	fmt.Println("  ╚════██║██║██╔═██╗ ██╔══██║╚════██║")
	fmt.Println("  ███████║██║██║  ██╗██║  ██║███████║")
	fmt.Println("  ╚══════╝╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝")
	fmt.Println()
	fmt.Println("  Student Cash Information Application")
	fmt.Println("  Telkom University · IF-49-INT · 2026")
	fmt.Println()
}
func printGoodbye() {
	fmt.Println()
	fmt.Println("  👋 Thank you for using SIKAS!")
	fmt.Println("  💙 Developed with ❤️  by Omar & Gama")
	fmt.Println("  🏫 Telkom University · IF-49-INT · 2026")
	fmt.Println()
}

//=======================MAIN FUNCTION========================
//============================================================

func main() {
	printSplash()
	var students StudentArray // fixed-size array of students
	var n int = 0             // number of students currently stored
	running := true

	for running { // repeat until user chooses 0 — repeat-until pattern
		printMenu()
		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		if choice == 1 {
			addStudent(&students, &n)
		} else if choice == 2 {
			editStudent(&students, n)
		} else if choice == 3 {
			deleteStudent(&students, &n)
		} else if choice == 4 {
			recordPayment(&students, n)
		} else if choice == 5 {
			fmt.Println("----------------------------------------- All Students -----------------------------------------")
			if n == 0 {
				fmt.Println("No students recorded yet.")
			} else {
				printAllStudents(students, n)
			}
		} else if choice == 6 {
			seqSearchUnpaid(students, n)
		} else if choice == 7 {
			binarySearchUnpaid(students, n)
		} else if choice == 8 {
			sortMenu(&students, n)
		} else if choice == 9 {
			showStatistics(students, n)
		} else if choice == 0 {
			running = false
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}

	printGoodbye()
}
