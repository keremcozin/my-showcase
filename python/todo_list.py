# A simple To-Do List Program written in Python

tasks = []
completed_tasks = []

def display_menu():
    print("Todo List Menu:")
    print("1. Add Task")
    print("2. View Pending Tasks")
    print("3. View Completed Tasks")
    print("4. Mark Task as Completed")
    print("5. Exit")

def add_task():
    task_name = input("Enter the task name: ")
    tasks.append(task_name)
    print("Task added successfully!")

def view_pending_tasks():
    if not tasks:
        print("No pending tasks.")
    else:
        print("Pending Tasks:")
        for index, task in enumerate(tasks, start=1):
            print(f"{index}. {task}")

def view_completed_tasks():
    if not completed_tasks:
        print("No completed tasks.")
    else:
        print("Completed Tasks:")
        for index, task in enumerate(completed_tasks, start=1):
            print(f"{index}. {task}")

def mark_completed():
    view_pending_tasks()
    if not tasks:
        return

    task_index = int(input("Enter the number of the task to mark as completed: ")) - 1
    if 0 <= task_index < len(tasks):
        completed_task = tasks.pop(task_index)
        completed_tasks.append(completed_task)
        print(f"Task '{completed_task}' marked as completed!")
    else:
        print("Invalid task number.")

def main():
    while True:
        display_menu()
        choice = input("Enter your choice (1-5): ")

        if choice == "1":
            add_task()
        elif choice == "2":
            view_pending_tasks()
        elif choice == "3":
            view_completed_tasks()
        elif choice == "4":
            mark_completed()
        elif choice == "5":
            print("Exiting the program.")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()