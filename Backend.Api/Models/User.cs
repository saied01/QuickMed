using System;
using System.ComponentModel.DataAnnotations;
using System.Collections.Generic;


namespace QuickMed.Models
{
    public class User
    {
        [Key]
        public int Id { get; set; }

        [Required]
        [MaxLength(50)]
        public string UserName { get; set; }

        [Required]
        [EmailAddress]
        public string Email { get; set; }

        [Required]
        public string PasswordHash { get; set; }

        public DateTime CreateDate { get; set; } = DateTime.UtcNow;

        public List<Reservation> Reservations { get; set; } = new();

    }
}
