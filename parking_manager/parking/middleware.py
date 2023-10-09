from parking import models
from django.shortcuts import get_object_or_404
from django.core.cache import cache

class ManagerMiddleware:
    """Check that manager set correct auth header"""
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):

        if "X-Manager-Login" in request.headers:
            manager = cache.get(f'manager_{request.headers["X-Manager-Login"]}')
            if not manager:
                manager = get_object_or_404(models.Manager, login=request.headers["X-Manager-Login"])
            request.manager = manager

        return self.get_response(request)
    
class ParkingMiddleware:
    """Add default parking to request"""
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):

        parking = cache.get("default_parking")
        if not parking:
            try:
                parking = models.Parking.objects.get(pk=1)
            except models.Parking.DoesNotExist:
                parking = models.Parking.objects.create()

            cache.set("default_parking", parking)
        
        request.parking = parking

        return self.get_response(request)