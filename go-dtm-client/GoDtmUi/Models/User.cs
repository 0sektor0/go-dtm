using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class User
    {
        [JsonProperty("id")]
        public int Id { get; set; }

        [JsonProperty("login")]
        public string Login { get; set; }
        
        public string Password { get; set; }

        [JsonProperty("picture")]
        public string Picture { get; set; }

        [JsonProperty("isAdmin")]
        public bool IsAdmin { get; set; }
        
        
        public static User FromJson(string json) => JsonConvert.DeserializeObject<User>(json, Converter.Settings);
    }
}
