using System.Collections.ObjectModel;
using DynamicData;
using GoDtmUI.Models;
using GoDtmUI.Api;
using ReactiveUI;



namespace GoDtmUI.ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        public TasksFilter Filter { get; } = new TasksFilter();
        public Client Client { get; } = new Client("http://194.87.144.249:8080");
        public ObservableCollection<Task> Tasks { get; } = new ObservableCollection<Task>();
        public ObservableCollection<Status> TaskTypes { get; } = new ObservableCollection<Status>();

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


        public MainWindowViewModel()
        {
            Client.OnApiError += error => Error = $"{error}\n{Error}";
            Client.OnAuthFailed += () => { IsAuthorized = false; };
            Client.OnAuthSuccess += () => { IsAuthorized = true; };
        }

        private void GetStatuses()
        {
            var types = Client.GetTypes();
            
            if (types != null)
            {
                TaskTypes.Clear();
                TaskTypes.AddRange(types.Statuses);
            }
        }

        public void OpenTaskInfo()
        {
            OpenedTask = Client.GetTask(TaskId);
        }

        public void SignIn()
        {
            if (Client.Auth(Login, Password))
            {
                GetTasks();
                GetStatuses();
            }
        }

        public void GetTasks()
        {
            var tasksInfo = Client.GetTasks(Filter);
            
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
    }
}
