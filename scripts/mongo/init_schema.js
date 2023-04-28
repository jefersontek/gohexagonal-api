db = db.getSiblingDB('account');

db.createCollection('accounts');

db.accounts.insertMany([
    {
        id: '1',
        account_id: '0001-1',
        name: 'User test',
        owner: 'customer',
        balance: 0.0,
        created_at: new Date()
    }
]);