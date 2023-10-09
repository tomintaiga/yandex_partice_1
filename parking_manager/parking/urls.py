from django.urls import path
from parking import api

urlpatterns = [
    path("managers/register", api.RegisterManager.as_view(), name="manager_register"),
]