# Security for the API
## Backend
- Add a login, so not everybody can participate
    - alternative: User have to give all the required infos like first- lastname, email, payment method
    - professors need a login, so only they can add / update their courses
- Only allow certain fields to be updated
- Validation on creation & update
- Escape input fields
- Go may be quick, but a DDoS protection by delaying every request with a second timeout can reduce DB stress
## Frontend
- Validation
