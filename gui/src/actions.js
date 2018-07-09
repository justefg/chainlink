import {
  getAccountBalance,
  getBridges,
  getConfiguration,
  getJobs,
  getJobSpec,
  getJobSpecRuns,
  getJobSpecRun
} from 'api'

export const REQUEST_JOBS = 'REQUEST_JOBS'
export const RECEIVE_JOBS_SUCCESS = 'RECEIVE_JOBS_SUCCESS'
export const RECEIVE_JOBS_ERROR = 'RECEIVE_JOBS_ERROR'

const requestJobs = () => ({ type: REQUEST_JOBS })
const receiveJobsSuccess = (json) => {
  return {
    type: RECEIVE_JOBS_SUCCESS,
    count: json.meta.count,
    items: json.data.map((j) => (
      {
        id: j.id,
        createdAt: j.attributes.createdAt,
        initiators: j.attributes.initiators
      }
    ))
  }
}
const receiveJobsNetworkError = (error) => {
  return {
    type: RECEIVE_JOBS_ERROR,
    error: error,
    networkError: true
  }
}

export const fetchJobs = (page, size) => {
  return dispatch => {
    dispatch(requestJobs())
    return getJobs(page, size)
      .then(json => dispatch(receiveJobsSuccess(json)))
      .catch(error => dispatch(receiveJobsNetworkError(error)))
  }
}

export const REQUEST_ACCOUNT_BALANCE = 'REQUEST_ACCOUNT_BALANCE'
export const RECEIVE_ACCOUNT_BALANCE_SUCCESS = 'RECEIVE_ACCOUNT_BALANCE_SUCCESS'
export const RECEIVE_ACCOUNT_BALANCE_ERROR = 'RECEIVE_ACCOUNT_BALANCE_ERROR'

const requestAccountBalance = () => ({ type: REQUEST_ACCOUNT_BALANCE })
const receiveAccountBalance = (json) => {
  return {
    type: RECEIVE_ACCOUNT_BALANCE_SUCCESS,
    eth: json.data.attributes.ethBalance,
    link: json.data.attributes.linkBalance
  }
}
const receiveAccountBalanceNetworkError = (error) => {
  return {
    type: RECEIVE_ACCOUNT_BALANCE_ERROR,
    error: error,
    networkError: true
  }
}

export const fetchAccountBalance = () => {
  return dispatch => {
    dispatch(requestAccountBalance())
    return getAccountBalance()
      .then(json => dispatch(receiveAccountBalance(json)))
      .catch(error => dispatch(receiveAccountBalanceNetworkError(error)))
  }
}

export const REQUEST_JOB_SPEC = 'REQUEST_JOB_SPEC'
export const RECEIVE_JOB_SPEC_SUCCESS = 'RECEIVE_JOB_SPEC_SUCCESS'
export const RECEIVE_JOB_SPEC_ERROR = 'RECEIVE_JOB_SPEC_ERROR'

const requestJobSpec = () => ({ type: REQUEST_JOB_SPEC })
const receiveJobSpecSuccess = (json) => {
  return {
    type: RECEIVE_JOB_SPEC_SUCCESS,
    item: json.data.attributes
  }
}
const receiveJobSpecNetworkError = (error) => {
  return {
    type: RECEIVE_JOB_SPEC_ERROR,
    networkError: true,
    error: error
  }
}

export const fetchJobSpec = (id) => {
  return dispatch => {
    dispatch(requestJobSpec())
    return getJobSpec(id)
      .then(json => dispatch(receiveJobSpecSuccess(json)))
      .catch(error => dispatch(receiveJobSpecNetworkError(error)))
  }
}

export const RECEIVE_JOB_SPEC_RUNS_SUCCESS = 'RECEIVE_JOB_SPEC_RUNS_SUCCESS'
export const RECEIVE_JOB_SPEC_RUNS_ERROR = 'RECEIVE_JOB_SPEC_RUNS_ERROR'

const receiveJobSpecRunsSuccess = (json) => {
  return {
    type: RECEIVE_JOB_SPEC_RUNS_SUCCESS,
    items: json.data.map(j => j.attributes),
    runsCount: json.meta.count
  }
}
const receiveJobSpecRunsNetworkError = (error) => {
  return {
    type: RECEIVE_JOB_SPEC_RUNS_ERROR,
    error: error,
    networkError: true
  }
}

export const fetchJobSpecRuns = (id, page, size) => {
  return dispatch => {
    return getJobSpecRuns(id, page, size)
      .then(json => dispatch(receiveJobSpecRunsSuccess(json)))
      .catch(error => dispatch(receiveJobSpecRunsNetworkError(error)))
  }
}

export const REQUEST_JOB_SPEC_RUN = 'REQUEST_JOB_SPEC_RUN'
export const RECEIVE_JOB_SPEC_RUN_SUCCESS = 'RECEIVE_JOB_SPEC_RUN_SUCCESS'
export const RECEIVE_JOB_SPEC_RUN_ERROR = 'RECEIVE_JOB_SPEC_RUN_ERROR'

const requestJobSpecRun = () => ({ type: REQUEST_JOB_SPEC_RUN })
const receiveJobSpecRunSuccess = (json) => {
  return {
    type: RECEIVE_JOB_SPEC_RUN_SUCCESS,
    item: json.data.attributes
  }
}
const receiveJobSpecRunNetworkError = (error) => {
  return {
    type: RECEIVE_JOB_SPEC_RUN_ERROR,
    error: error,
    networkError: true
  }
}
export const fetchJobSpecRun = (id) => {
  return dispatch => {
    dispatch(requestJobSpecRun())
    return getJobSpecRun(id)
      .then(json => dispatch(receiveJobSpecRunSuccess(json)))
      .catch(error => dispatch(receiveJobSpecRunNetworkError(error)))
  }
}

export const REQUEST_CONFIGURATION = 'REQUEST_CONFIGURATION'
export const RECEIVE_CONFIGURATION_SUCCESS = 'RECEIVE_CONFIGURATION_SUCCESS'
export const RECEIVE_CONFIGURATION_ERROR = 'RECEIVE_CONFIGURATION_ERROR'

const requestConfiguration = () => ({ type: REQUEST_CONFIGURATION })
const receiveConfiguration = (json) => {
  return {
    type: RECEIVE_CONFIGURATION_SUCCESS,
    config: json.data.attributes
  }
}
const receiveConfigurationNetworkError = (error) => {
  return {
    type: RECEIVE_CONFIGURATION_ERROR,
    error: error,
    networkError: true
  }
}
export const fetchConfiguration = () => {
  return dispatch => {
    dispatch(requestConfiguration())
    return getConfiguration()
      .then(json => dispatch(receiveConfiguration(json)))
      .catch(error => dispatch(receiveConfigurationNetworkError(error)))
  }
}

export const REQUEST_BRIDGES = 'REQUEST_BRIDGES'
export const RECEIVE_BRIDGES_SUCCESS = 'RECEIVE_BRIDGES_SUCCESS'
export const RECEIVE_BRIDGES_ERROR = 'RECEIVE_BRIDGES_ERROR'

const requestBridges = () => ({ type: REQUEST_BRIDGES })
const receiveBridgesSuccess = (json) => {
  return {
    type: RECEIVE_BRIDGES_SUCCESS,
    count: json.meta.count,
    items: json.data.map(b => b.attributes)
  }
}
const receiveBridgesNetworkError = (error) => {
  return {
    type: RECEIVE_BRIDGES_ERROR,
    error: error,
    networkError: true
  }
}

export const fetchBridges = (page, size) => {
  return dispatch => {
    dispatch(requestBridges())
    return getBridges(page, size)
      .then(json => dispatch(receiveBridgesSuccess(json)))
      .catch(error => dispatch(receiveBridgesNetworkError(error)))
  }
}
