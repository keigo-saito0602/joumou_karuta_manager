CREATE INDEX idx_event_scores_rank ON event_scores (
    score DESC,
    cards_taken DESC,
    fault_count ASC,
    created_at ASC
);
