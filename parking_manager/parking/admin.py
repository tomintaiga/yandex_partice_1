from django.contrib import admin
from parking import models

admin.site.register(models.Parking)
admin.site.register(models.Manager)
admin.site.register(models.Employee)
admin.site.register(models.ParkingSlot)
admin.site.register(models.Booking)
