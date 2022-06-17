Rails.application.routes.draw do
  resources :contacts
  resources :users, param: :_username
  resources :facts
  post '/auth/login', to: 'authentication#login'
  get '/*a', to: 'application#not_found'
end
