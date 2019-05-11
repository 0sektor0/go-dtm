using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using DynamicData;
using GoDtmUI.Models;
using GoDtmUI.Api;
using GoDtmUI.Utils;
using ReactiveUI;
using Task = GoDtmUI.Models.Task;



namespace GoDtmUI.ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        public Client Client { get; private set; }
        public TasksFilter Filter { get; } = new TasksFilter();
        public ObservableCollection<User> Users { get; } = new ObservableCollection<User>();
        public ObservableCollection<Task> Tasks { get; } = new ObservableCollection<Task>();
        public ObservableCollection<Status> TaskTypes { get; } = new ObservableCollection<Status>();
        public string Server { get; set; } = "http://192.168.1.4:8080";
        public string NewTaskTitle { get; set; }
        public string NewTaskText { get; set; }
        public int TaskToDeleteId { get; set; }

        private Dictionary<string, User> _users;
        private Dictionary<string, Status> _taskStatuses;
        private Task _openedTask;
        private bool _isAuthorized;
        private string _password;
        private string _login;
        private string _error;

        public bool IsAuthorized
        {
            get => _isAuthorized;
            set => this.RaiseAndSetIfChanged(ref _isAuthorized, value);
        }

        public string Login
        {
            get => _login; 
            set => this.RaiseAndSetIfChanged(ref _login, value); 
        }
        
        public string Password
        {
            get  => _password; 
            set  => this.RaiseAndSetIfChanged(ref _password, value); 
        }
        
        public string Error
        {
            get => _error; 
            set => this.RaiseAndSetIfChanged(ref _error, value); 
        }

        public Task OpenedTask
        {
            get => _openedTask; 
            set => this.RaiseAndSetIfChanged(ref _openedTask, value); 
        }

        public int TaskId { get; set; }

        public string CommentText { get; set; }

        
        public void SignIn()
        {
            Client = new Client(Server);
            Client.OnApiError += error => Error = $"{error}\n{Error}";
            Client.OnAuthFailed += () => { IsAuthorized = false; };
            Client.OnAuthSuccess += () => { IsAuthorized = true; };
            
            if (Client.Auth(Login, Password))
            {
                GetTasks();
                GetStatuses();
                GetUsers();
            }
        }
        
        public void LogOut()
        {
            Client.LogOut();
        }

        public void GetUsers()
        {
            Users.Clear();
            Users.AddRange(Client.GetUsers(int.MaxValue, 0)?.Users);
            
            if(Users != null)
                _users = Users.ToDictionary(u => u.Login, u => u);
        }        

        public void GetStatuses()
        {
            var types = Client.GetTypes();
            
            if (types != null)
            {
                TaskTypes.Clear();
                TaskTypes.AddRange(types.Statuses);

                _taskStatuses = types.Statuses.ToDictionary(t => t.Name, t => t);
            }
        }

        public void OpenTaskInfo()
        {
            OpenedTask = Client.GetTask(TaskId);
        }

        public void GetTasks()
        {
            var filter = new TasksFilter
            {
                Limit = Filter.Limit,
                Offset = Filter.Offset
            };

            var status = _taskStatuses.GetValueAlways(Filter.Status.Name);
            if (status != null)
                filter.Status = status;

            var asignee = _users.GetValueAlways(Filter.Asignee.Login);
            if (asignee != null)
                filter.Asignee = asignee;
                    
            var creator = _users.GetValueAlways(Filter.Creator.Login);
            if (creator != null)
                filter.Creator = creator;

            if (Filter.StartDate != default)
                filter.StartDate = Filter.StartDate;

            if (Filter.EndDate != default)
                filter.EndDate = Filter.EndDate;
            
            var tasksInfo = Client.GetTasks(filter);
            if (tasksInfo != null)
            {
                Tasks.Clear();
                Tasks.AddRange(tasksInfo.Tasks);
            }
        }
        
        public void AddComment()
        {
            if (OpenedTask == null) return;

            Client.CreateComment(OpenedTask.Id, CommentText);
            OpenedTask = Client.GetTask(OpenedTask.Id);
        }

        public void UpdateTask()
        {
            if (_users.ContainsKey(OpenedTask.Asignee.Login))
                OpenedTask.Asignee = _users[OpenedTask.Asignee.Login];

            if (_taskStatuses.ContainsKey(OpenedTask.Status.Name))
                OpenedTask.Status = _taskStatuses[OpenedTask.Status.Name];

            if (Client.UpdateTask(OpenedTask))
            {
                GetTasks();
                OpenTaskInfo();
            }
        }

        public void CreateTask()
        {
            var newTask = Client.CreateTask(1, NewTaskTitle, NewTaskText);
            if(newTask != null)
                Tasks.Add(newTask);
        }

        public void DeleteTask()
        {
            if (Client.DeleteTask(TaskToDeleteId))
            {
                var taskToDelete = Tasks.Where(t => t.Id == TaskToDeleteId).FirstOrDefault();
                Tasks.Remove(taskToDelete);
            }
        }
    }
}
