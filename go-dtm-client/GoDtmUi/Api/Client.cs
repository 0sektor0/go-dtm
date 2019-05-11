using System;
using System.Net;
using System.Net.Http;
using System.Text;
using GoDtmUI.Models;



namespace GoDtmUI.Api
{
    public class Client
    {
        private string _host;
        private HttpClient _http;
        private Session _session;

        public event Action OnAuthFailed;
        public event Action OnAuthSuccess;
        public event Action<string> OnApiError;
        
        public User User => _session?.User;


        public Client(string host)
        {
            _host = host;
            _http = new HttpClient();
        }

        private T SendApiRequest<T>(string content, string apiMethod, Func<string, T> parser) where T : class
        {
            var request = new HttpRequestMessage(HttpMethod.Post, $"{_host}/{apiMethod}");
            request.Content = new StringContent(content, Encoding.UTF8, "application/x-www-form-urlencoded");

            HttpResponseMessage response;
            try
            {
                response = _http.SendAsync(request).Result;
            }
            catch (Exception e)
            {
                LogOut();
                OnApiError?.Invoke(e.Message);
                return null;
            }
            
            var responseBody = response.Content.ReadAsStringAsync().Result;

            if (response.StatusCode != HttpStatusCode.OK)
            {
                OnApiError?.Invoke(responseBody);
                return null;
            }

            try
            {
                var data = parser(responseBody);
                return data;
            }
            catch (Exception ex)
            {
                OnApiError?.Invoke(ex.Message);
                return null;
            }
        }
        
        public bool Auth(string login, string password)
        {
            _session = SendApiRequest($"login={login}&password={password}", "auth", Session.FromJson);
            if (_session == null)
            {
                OnAuthFailed?.Invoke();
                return false;
            }
            
            _session.User.Password = password;
            OnAuthSuccess?.Invoke();
            
            return true;
        }

        private bool Auth()
        {
            return Auth(User.Login, User.Password);
        }

        public void LogOut()
        {
            OnAuthFailed?.Invoke();
        }

        public User SignUp(string login, string password)
        {
            var user = SendApiRequest($"login={login}&password={password}","signup",User.FromJson);
            return user;
        }
        
        private void CheckAuthState()
        {
            if (_session == null)
            {
                throw new AuthException();
            }
            else if (_session.IsExpired() && !Auth())
            {
                throw new AuthException();
            }
        }

        public TasksInfo GetTasks(TasksFilter filter)
        {
            CheckAuthState();

            var tasks = SendApiRequest(
                $"token={_session.Token}&offset={filter.Offset}&limit={filter.Limit}&filterType={filter.Type}&filterParam={filter.Param}", 
                "tasks/get",
                TasksInfo.FromJson);
            
            return tasks;
        }

        public Task GetTask(int id)
        {
            CheckAuthState();

            var task = SendApiRequest(
                $"token={_session.Token}&id={id}", 
                "task/get",
                Task.FromJson);
            
            return task;
        }

        public bool DeleteTask(int id)
        {
            CheckAuthState();
            
            var result =  SendApiRequest(
                $"token={_session.Token}&id={id}", 
                "tasks/delete",
                FromJson);

            return result != null;
        }

        public UsersInfo GetUsers(int limit, int offset)
        {
            CheckAuthState();

            var users = SendApiRequest(
                $"token={_session.Token}&offset={offset}&limit={limit}", 
                "users/get",
                UsersInfo.FromJson);
            
            return users;
        }

        public User GetUser(int id)
        {
            return GetUser(id, "");
        }
        
        public User GetUser(string name)
        {
            return GetUser(0, name);
        }
        
        private User GetUser(int id, string login)
        {
            CheckAuthState();

            var user = SendApiRequest(
                $"token={_session.Token}&id={id}&login={login}", 
                "user/get",
                User.FromJson);
            
            return user;
        }

        public bool UpdateTask(Task task)
        {
            CheckAuthState();

            var taskJson = Converter.ToJson(task);
            var result =  SendApiRequest(
                $"token={_session.Token}&id={task.Id}&taskUpdate={taskJson}", 
                "tasks/update",
                FromJson);
            
            return result != null;
        }
        
        public Task CreateTask(int taskType, string title, string text)
        {
            CheckAuthState();

            var task =  SendApiRequest(
                $"token={_session.Token}&id={taskType}&text={text}&title={title}", 
                "tasks/create",
                Task.FromJson);

            return task;
        }

        public bool CreateComment(int taskId, string text)
        {
            CheckAuthState();

            var result =  SendApiRequest(
                $"token={_session.Token}&id={taskId}&commentText={text}", 
                "comments/add",
                Task.FromJson);
            
            return result != null;
        }
        
        public bool UpdateComment(int commentId, string text)
        {
            CheckAuthState();

            var result =  SendApiRequest(
                $"token={_session.Token}&id={commentId}&commentText={text}", 
                "comments/update",
                Task.FromJson);
            
            return result != null;
        }
        
        public bool DeleteComment(int commentId)
        {
            CheckAuthState();

            var result =  SendApiRequest(
                $"token={_session.Token}&id={commentId}", 
                "comments/delete",
                Task.FromJson);
            
            return result != null;
        }
        
        public TypesInfo GetTypes()
        {
            CheckAuthState();

            var types =  SendApiRequest(
                $"token={_session.Token}", 
                "statuses/get",
                TypesInfo.FromJson);
            
            return types;
        }
        
        private string FromJson(string json)
        {
            return json;
        }
    }
}