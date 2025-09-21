using System;
using System.ComponentModel.DataAnnotations;
using System.Collections.Generic;
using Microsoft.AspNetCore.Identity;

namespace QuickMed.Models
{
    public enum UserType { Patient, Provider }

    public class User : IdentityUser
    {

        [Required]
        public string FirstName { get; set; }
        [Required]
        public string LastName { get; set; }

        public UserType Type { get; set; }
        // Solo aplica si es doctor
        public string? Specialty { get; set; }
        public string? LicenseNumber { get; set; }
        public bool? IsVerified { get; set; }

        // Relación con reservas (solo para pacientes)
        public List<Reservation>? Reservations { get; set; }
    }
}
