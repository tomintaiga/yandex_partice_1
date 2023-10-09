from django.db import models

class Parking(models.Model):
    scheme = models.FileField(upload_to="scheme")

class Manager(models.Model):
    login = models.CharField(max_length=128)
    parking = models.ForeignKey(Parking, on_delete=models.CASCADE)

    def __str__(self) -> str:
        return self.login

class Employee(models.Model):
    login = models.CharField(max_length=128)
    balance = models.IntegerField()
    month_limit = models.IntegerField()
    manager = models.ForeignKey(Manager, on_delete=models.CASCADE)

    def __str__(self) -> str:
        return self.login


class ParkingSlot(models.Model):
    name = models.CharField(max_length=128)
    parking = models.ForeignKey(Parking, on_delete=models.CASCADE)

class Booking(models.Model):

    class State(models.TextChoices):
        CANCELED = "C", "Canceled"
        BOOKED = "B", "Booked"

    slot = models.ForeignKey(ParkingSlot, on_delete=models.CASCADE)
    date = models.DateField()
    car_plate = models.CharField(max_length=128)
    status = models.CharField(max_length=10, choices=State.choices, default=State.BOOKED)
    employee = models.ForeignKey(Employee, on_delete=models.CASCADE)
    