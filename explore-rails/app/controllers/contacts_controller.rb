class ContactsController < ApplicationController
    before_action :authorize_request

    def index 
        @contacts = Contact.all
        render json: @contacts
    end

    def show
        @contact = Contact.find(params[:id])
        render json: @contact
    end

    def create
        @contact = Contact.new(contact_params)
        if @contact.save
            render json: @contact
        else 
            render error: { error: 'Unable to create Contact.' }, status: 400
        end
    end

    def update 
        @contact = Contact.find(params[:id])
        if @contact
            @contact.update(contact_params)
            render json: { message: 'Contact successfully updated.' }, status: 200
        else 
            render error: { error: 'Unable to create Contact.' }, status: 400
        end
    end

    def destroy
        @contact = Contact.find(params[:id])
        if @contact
            @contact.destroy
            render json: { message: 'Contact successfully deleted.' }, status: 200
        else 
            render error: { error: 'Unable to create Contact.' }, status: 400
        end
    end

    private 

    def contact_params
        params.require(:contact).permit(:name, :email, :phone, :website, :address)
    end

end