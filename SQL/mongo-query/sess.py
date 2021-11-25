from pymongo import 


def run_transaction_with_retry(txn_func, session):
    while True:
        try:
            txn_func(session)  # performs transaction
            break
        except (ConnectionFailure, OperationFailure) as exc:
            # If transient error, retry the whole transaction
            if exc.has_error_label("TransientTransactionError"):
                print("TransientTransactionError, retrying "
                      "transaction ...")
                continue
            else:
                raise

def commit_with_retry(session):
    while True:
        try:
            # Commit uses write concern set at transaction start.
            session.commit_transaction()
            print("Transaction committed.")
            break
        except (ConnectionFailure, OperationFailure) as exc:
            # Can retry commit
            if exc.has_error_label("UnknownTransactionCommitResult"):
                print("UnknownTransactionCommitResult, retrying "
                      "commit operation ...")
                continue
            else:
                print("Error during commit ...")
                raise

# Updates two collections in a transactions

def update_employee_info(session):
    employees_coll = session.client.hr.employees
    events_coll = session.client.reporting.events

    with session.start_transaction(
            read_concern=ReadConcern("snapshot"),
            write_concern=WriteConcern(w="majority"),
            read_preference=ReadPreference.PRIMARY):
        employees_coll.update_one(
            {"employee": 3}, {"$set": {"status": "Inactive"}},
            session=session)
        events_coll.insert_one(
            {"employee": 3, "status": {
                "new": "Inactive", "old": "Active"}},
            session=session)

        commit_with_retry(session)

# Start a session.
with client.start_session() as session:
    try:
        run_transaction_with_retry(update_employee_info, session)
    except Exception as exc:
        # Do something with error.
        raise
