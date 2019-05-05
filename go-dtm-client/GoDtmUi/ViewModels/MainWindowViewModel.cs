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
        
        

        public MainWindowViewModel()
        {
            Client.OnApiError += error => Error = $"{error}\n{Error}";
            Client.OnAuthFailed += () => { IsAuthorized = false; };
            Client.OnAuthSuccess += () => { IsAuthorized = true; };
        }

        public void OpenTaskInfo(int id)
        {
            var a = "";
        }

        public void SignIn()
        {
            if(Client.Auth(Login, Password))
                GetTasks();
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
    }
}
