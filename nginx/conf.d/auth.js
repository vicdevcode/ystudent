function introspectAccessToken(r) {
  r.subrequest("/_auth_send_request", function(reply) {
    if (reply.status === 200) r.return(204);
    else r.return(401);
  });
}

export default { introspectAccessToken };
