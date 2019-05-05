using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public partial class Comment
    {
        [JsonProperty("id")]
        public int Id { get; set; }

        [JsonProperty("developer")]
        public User Developer { get; set; }

        [JsonProperty("developerId")]
        public int DeveloperId { get; set; }

        [JsonProperty("text")]
        public string Text { get; set; }
        
        
        public static Comment FromJson(string json) => JsonConvert.DeserializeObject<Comment>(json, Converter.Settings);
    }
}