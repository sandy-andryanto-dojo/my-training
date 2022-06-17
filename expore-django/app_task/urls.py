from django.urls import path
from . import views

urlpatterns = [
	path('tasks/list/', views.taskList, name="tasks.list"),
	path('tasks/detail/<str:pk>/', views.taskDetail, name="tasks.detail"),
	path('tasks/create/', views.taskCreate, name="tasks.create"),
	path('tasks/update/<str:pk>/', views.taskUpdate, name="task.update"),
	path('tasks/delete/<str:pk>/', views.taskDelete, name="task.delete"),
]
