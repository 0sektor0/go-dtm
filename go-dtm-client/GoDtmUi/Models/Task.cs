using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class Task
    {
        [JsonProperty("id")]
        public int Id { get; set; }

        [JsonProperty("creator")]
        public User Creator { get; set; }

        [JsonProperty("asignee")]
        public User Asignee { get; set; }

        [JsonProperty("type")]
        public Status Type { get; set; }

        [JsonProperty("status")]
        public Status Status { get; set; }

        [JsonProperty("title")]
        public string Title { get; set; }

        [JsonProperty("text")]
        public string Text { get; set; }

        [JsonProperty("creationDate")]
        public int CreationDate { get; set; }

        [JsonProperty("startDate")]
        public int StartDate { get; set; }

        [JsonProperty("endDate")]
        public int EndDate { get; set; }

        [JsonProperty("updDateTime")]
        public int UpdDateTime { get; set; }

        [JsonProperty("attachments")]
        public Attachment[] Attachments { get; set; }

        [JsonProperty("comments")]
        public Comment[] Comments { get; set; }
        
        
        public static Task FromJson(string json) => JsonConvert.DeserializeObject<Task>(json, Converter.Settings);
    }
}