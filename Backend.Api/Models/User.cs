using System;
using System.ComponentModel.DataAnnotations;
using System.Collections.Generic;


namespace QuickMed.Models
{

    public enum UserType { Patient, Doctor }

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

        public string FirstName { get; set; }

        public string LastName { get; set; }

        [Required]
        public string PasswordHash { get; set; }

        public DateTime CreateDate { get; set; } = DateTime.UtcNow;

        public List<Reservation> Reservations { get; set; } = new();

        // Only applies if user is doctor
        public string? Specialty { get; set; }
        public string? LicenseNumber { get; set; }
        public bool? IsVerified { get; set; } = false;

    }

}
