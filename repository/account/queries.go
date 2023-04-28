package account

const findByID = `
select id, account_id, name, owner, balance, created_at
 from accounts
where account_id = ?
`
