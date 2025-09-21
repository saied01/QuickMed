using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using QuickMed.Models;

namespace QuickMed.Data
{
    public class AppDbContext : IdentityDbContext<ApplicationUser>
    {
        public AppDbContext(DbContextOptions<AppDbContext> options)
            : base(options) { }

        public DbSet<Resource> Resources { get; set; }
        public DbSet<Reservation> Reservations { get; set; }

        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);

            // index for available quick reservations
            builder.Entity<Reservation>()
                .HasIndex(r => new { r.ResourceId, r.StartUtc });

            // Aquí podrías agregar constraints de PostgreSQL para evitar solapamiento de reservas
        }
    }
}
