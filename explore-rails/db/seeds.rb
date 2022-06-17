# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)
require 'faker'

for i in 1..100 do
    user = User.new(
        :username  => Faker::Internet.username,
        :email => Faker::Internet.safe_email,
        :password=> "password",
        :is_active => true
    )
    user.save!
end

for i in 1..100 do
    contact = Contact.new(
        :name  => Faker::Name.unique.name,
        :email => Faker::Internet.safe_email,
        :phone => Faker::PhoneNumber.phone_number,
        :website  => Faker::Internet.url,
        :address => Faker::Lorem.paragraph
    )
    contact.save!
end