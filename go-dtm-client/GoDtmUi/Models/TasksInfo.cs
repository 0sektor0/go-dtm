using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class TasksInfo
    {
        [JsonProperty("tasks")]
        public Task[] Tasks { get; set; }
        
        
        public static TasksInfo FromJson(string json) => JsonConvert.DeserializeObject<TasksInfo>(json, Converter.Settings);
    }
}