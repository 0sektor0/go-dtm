using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class UsersInfo
    {
        [JsonProperty("users")]
        public User[] Users { get; set; }
        
        
        public static UsersInfo FromJson(string json) => JsonConvert.DeserializeObject<UsersInfo>(json, Converter.Settings);
    }
}