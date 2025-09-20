using System.ComponentModel.DataAnnotations;
using System.Collections.Generic;



namespace QuickMed.Models
{
    public class Resource
    {
        [Key]
        public int Id { get; set; }
        [Required]
        public string Name { get; set; }
        [Required]
        public string Type { get; set; } // room, doctor, etc...

        public int Capacity { get; set; }

        [Required]
        public string MetadataJson { get; set; }
    }


    public enum ReservationStatus
    {
        Booked,
        Cancelled,
        Finished
    }
    public class Reservation
    {
        [Key]
        public int Id { get; set; }

        [Required]
        public int ResourceId { get; set; }
        public Resource Resource { get; set; }

        [Required]
        public int UserId { get; set; }

        [Required]
        public DateTime StartUtc { get; set; }

        [Required]
        public DateTime EndUtc { get; set; }

        [Required]
        public ReservationStatus Status { get; set; } //booked, cancelled, finished, ...

        public DateTime CreatedAt { get; set; } = DateTime.UtcNow;

    }
}
